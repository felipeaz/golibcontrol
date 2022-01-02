package module

import (
	"encoding/json"
	"github.com/FelipeAz/golibcontrol/infra/http/client"
	"github.com/FelipeAz/golibcontrol/infra/http/request"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/pkg"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	studentErrors "github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository domain.Repository
	Log        logger.LogInterface
}

func NewStudentModule(repo domain.Repository, log logger.LogInterface) StudentModule {
	return StudentModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all students.
func (m StudentModule) Get() ([]students.Student, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one student by ID.
func (m StudentModule) Find(id string) (students.Student, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a student to the database.
func (m StudentModule) Create(student students.Student, accountHost, accountRoute, tokenName, tokenValue string) (*students.Student, *errors.ApiError) {
	accountId, apiError := m.createAccountOnAccountService(student, accountHost, accountRoute, tokenName, tokenValue)
	if apiError != nil {
		return nil, apiError
	}
	student.AccountId = accountId
	return m.Repository.Create(student)
}

// Update update an existent student.
func (m StudentModule) Update(id string, upStudent students.Student) *errors.ApiError {
	return m.Repository.Update(id, upStudent)
}

// Delete delete an existent student.
func (m StudentModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}

func (m StudentModule) createAccountOnAccountService(student students.Student, host, route, tokenName, tokenValue string) (uint, *errors.ApiError) {
	studentBody := pkg.ConvertStudentToStudentAccount(student)
	body, err := json.Marshal(studentBody)
	if err != nil {
		m.Log.Error(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToMarshalRequestBody,
			Error:   err.Error(),
		}
	}

	req := request.NewHttpRequest(client.NewHTTPClient(), host)
	b, err := req.PostWithHeader(route, body, tokenName, tokenValue)
	if err != nil {
		m.Log.Error(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToSendAccountCreationRequest,
			Error:   err.Error(),
		}
	}

	var resp domain.AccountResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		m.Log.Error(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToUnmarshalResponse,
			Error:   err.Error(),
		}
	}

	if resp.ID == 0 {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToCreateStudentAccount,
			Error:   string(b),
		}
	}
	return resp.ID, nil
}
