package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
	"strconv"
)

func ParseStringToId(id string) (uint, *errors.ApiError) {
	intVal, err := strconv.Atoi(id)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}
	return uint(intVal), nil
}
