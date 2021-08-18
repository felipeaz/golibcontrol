package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
)

func ConvertToStudentObj(obj interface{}) (model.Student, *errors.ApiError) {
	studentObj, ok := obj.(*model.Student)
	if !ok {
		return model.Student{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *studentObj, nil
}

func ConvertToSliceStudentObj(obj interface{}) ([]model.Student, *errors.ApiError) {
	studentObj, ok := obj.(*[]model.Student)
	if !ok {
		return nil, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *studentObj, nil
}

func ConvertStudentToStudentAccount(student model.Student) model.StudentAccount {
	return model.StudentAccount{
		Email:          student.Email,
		Password:       student.Password,
		FirstName:      student.GetFirstName(),
		LastName:       student.GetLastName(),
		Phone:          student.Phone,
		StudentAccount: true,
	}
}
