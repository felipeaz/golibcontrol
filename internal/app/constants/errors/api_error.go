package errors

// Default return login
const (
	// Rest Messages
	FailMessage                 = "Failed to fetch data"
	CreateFailMessage           = "Failed to create data"
	UpdateFailMessage           = "Failed to update data"
	DeleteFailMessage           = "Failed to delete data"
	JWTTokenCreationFailMessage = "Login Failed"

	// Input Association Messages
	FailedFieldsAssociationMessage = "Failed while associating fields from request"
)

// ApiError will be used on API Errors
type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
