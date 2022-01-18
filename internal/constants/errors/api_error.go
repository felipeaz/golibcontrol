package errors

// Default return login
const (
	FailMessage                       = "Failed to fetch data"
	CreateFailMessage                 = "Failed to create data"
	UpdateFailMessage                 = "Failed to update data"
	DeleteFailMessage                 = "Failed to delete data"
	FailedFieldsAssociationMessage    = "Failed while associating fields from request"
	FailedToParseObj                  = "Failed on object conversion"
	FailedToCreateBookCategoryMessage = "Failed to create book category"

	LendingNotAvailableError              = "lending not available"
	FailedToStoreAuthenticationKeyOnCache = "Failed to store authentication key on cache"
	FailedToGetAuthenticationOnCache      = "Failed to get authentication on cache"
	FailedToParseAuthenticationFromCache  = "Failed to parse authentication from cache"
	FailedToDeleteAuthenticationOnCache   = "Failed to delete authentication on cache"
	FailedToMarshal                       = "Failed to marshal data"
	FailedToCreateConsumer                = "Failed to create consumer"
	FailedToDeleteConsumer                = "Failed to delete consumer"
	FailedToRetrieveConsumerKey           = "Failed to retrieve consumer"
)

// ApiError will be used on API Errors
type ApiError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
