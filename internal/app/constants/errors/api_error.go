package errors

// Default return messages
const (
	// Rest Messages
	NotFoundMessage     = "Failed to fetch data"
	CreateFailedMessage = "Failed to create data"
	UpdateFailedMessage = "Failed to update data"
	DeleteFailedMessage = "Failed to delete data"

	// Input Association Messages
	FailedFieldsAssociationMessage = "Failed while associating fields from request"
)

// ApiError will be used on API Errors
type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
