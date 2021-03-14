package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// LendingModule process the request recieved from handler.
type LendingModule struct {
	Repository repository.LendingRepository
}
