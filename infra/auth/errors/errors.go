package errors

const (
	FailedToDeleteConsumer        = "Failed to delete consumer"
	FailedToCreateConsumer        = "Failed to create consumer"
	FailedToRetrieveConsumer      = "Failed to retrieve consumer"
	FailedToCleanConsumerKeys     = "Failed to clean consumer keys"
	FailedToRetrieveConsumerKeys  = "Failed to retrieve consumer keys"
	FailedToMarshalKeyRequestBody = "Failed to marshal key request body"
	FailedToMarshalConsumer       = "Failed to marshal consumer"
	FailedToUnmarshalConsumer     = "Failed to unmarshal consumer"
	JWTTokenCreationFailMessage   = "Failed to create JWT token"
)
