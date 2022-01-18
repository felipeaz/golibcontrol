package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToStudentObj(obj interface{}) (students.Student, *errors.ApiError) {
	data, ok := obj.(*students.Student)
	if !ok {
		return students.Student{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}

func ParseToSliceStudentObj(obj interface{}) ([]students.Student, *errors.ApiError) {
	if obj == nil {
		return []students.Student{}, nil
	}
	data, ok := obj.(*[]students.Student)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}

func ParseStudentToStudentAccount(student students.Student) students.Account {
	return students.Account{
		Email:          student.Email,
		Password:       student.Password,
		FirstName:      student.GetFirstName(),
		LastName:       student.GetLastName(),
		Phone:          student.Phone,
		StudentAccount: true,
	}
}
