package errors

// Default return login
const (
	// Messages
	FailMessage                       = "Failed to fetch data"
	CreateFailMessage                 = "Failed to create data"
	UpdateFailMessage                 = "Failed to update data"
	DeleteFailMessage                 = "Failed to delete data"
	JWTTokenCreationFailMessage       = "Login Failed"
	FailedFieldsAssociationMessage    = "Failed while associating fields from request"
	FailedToConvertObj                = "Failed on object conversion"
	AuthenticationFailMessage         = "Failed to authenticate the user"
	FailedToCreateBookCategoryMessage = "Failed to create book category"

	// Errors
	ItemNotFoundError        = "item not found"
	CategoryNotFoundError    = "category not found"
	StudentNotFoundError     = "studend not found"
	BookNotFoundError        = "book not found"
	LendingNotAvailableError = "lending not available"
)

// ApiError will be used on API Errors
type ApiError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
