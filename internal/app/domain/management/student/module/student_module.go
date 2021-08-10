package module

import (
	"encoding/json"
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	studentErrors "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model/converter"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/pkg/http/request"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository _interface.StudentRepositoryInterface
}

func NewStudentModule(repo _interface.StudentRepositoryInterface) StudentModule {
	return StudentModule{
		Repository: repo,
	}
}

// Get returns all students.
func (m StudentModule) Get() ([]model.Student, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one student by ID.
func (m StudentModule) Find(id string) (model.Student, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a student to the database.
func (m StudentModule) Create(student model.Student, accountHost, accountRoute, tokenName, tokenValue string) (string, *errors.ApiError) {
	accountId, apiError := m.createAccountOnAccountService(student, accountHost, accountRoute, tokenName, tokenValue)
	if apiError != nil {
		return "", apiError
	}
	student.AccountId = accountId
	return m.Repository.Create(student)
}

// Update update an existent student.
func (m StudentModule) Update(id string, upStudent model.Student) *errors.ApiError {
	return m.Repository.Update(id, upStudent)
}

// Delete delete an existent student.
func (m StudentModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}

func (m StudentModule) createAccountOnAccountService(student model.Student, host, route, tokenName, tokenValue string) (uint, *errors.ApiError) {
	var resp model.AccountResponse
	req := request.NewHttpRequest(host)
	studentBody := converter.ConvertStudentToStudentAccount(student)
	body, err := json.Marshal(studentBody)
	if err != nil {
		logger.LogError(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToMarshalRequestBody,
			Error:   err.Error(),
		}
	}

	b, err := req.PostRequestWithHeader(route, body, tokenName, tokenValue)
	if err != nil {
		logger.LogError(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: studentErrors.FailedToSendAccountCreationRequest,
			Error:   err.Error(),
		}
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		logger.LogError(err)
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
