package errors

// Default return login
const (
	// Messages
	FailMessage                       = "Failed to fetch data"
	CreateFailMessage                 = "Failed to create data"
	UpdateFailMessage                 = "Failed to update data"
	DeleteFailMessage                 = "Failed to delete data"
	FailedFieldsAssociationMessage    = "Failed while associating fields from request"
	FailedToConvertObj                = "Failed on object conversion"
	FailedToCreateBookCategoryMessage = "Failed to create book category"

	// Errors
	ItemNotFoundError        = "item not found"
	CategoryNotFoundError    = "category not found"
	LendingNotAvailableError = "lending not available"
)

// ApiError will be used on API Errors
type ApiError struct {
	Service string `json:"service"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
