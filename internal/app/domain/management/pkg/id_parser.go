package pkg

import (
	"net/http"
	"os"
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseStringToId(id string) (uint, *errors.ApiError) {
	intVal, err := strconv.Atoi(id)
	if err != nil {
		return 0, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}
	return uint(intVal), nil
}
