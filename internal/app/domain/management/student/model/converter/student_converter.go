package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
)

func ConvertToStudentObj(obj interface{}) (model.Student, *errors.ApiError) {
	studentObj, ok := obj.(*model.Student)
	if !ok {
		return model.Student{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *studentObj, nil
}

func ConvertToSliceStudentObj(obj interface{}) ([]model.Student, *errors.ApiError) {
	if obj == nil {
		return []model.Student{}, nil
	}
	studentObj, ok := obj.(*[]model.Student)
	if !ok {
		return nil, &errors.ApiError{
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
