package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ConvertToStudentObj(obj interface{}) (students.Student, *errors.ApiError) {
	studentObj, ok := obj.(*students.Student)
	if !ok {
		return students.Student{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *studentObj, nil
}

func ConvertToSliceStudentObj(obj interface{}) ([]students.Student, *errors.ApiError) {
	if obj == nil {
		return []students.Student{}, nil
	}
	studentObj, ok := obj.(*[]students.Student)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *studentObj, nil
}

func ConvertStudentToStudentAccount(student students.Student) students.Account {
	return students.Account{
		Email:          student.Email,
		Password:       student.Password,
		FirstName:      student.GetFirstName(),
		LastName:       student.GetLastName(),
		Phone:          student.Phone,
		StudentAccount: true,
	}
}
