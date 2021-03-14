package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository repository.StudentRepository
}
