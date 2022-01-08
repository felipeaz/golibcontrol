package model

type Consumer struct {
	Id        string      `json:"id,omitempty"`
	CustomId  string      `json:"custom_id,omitempty"`
	Username  string      `json:"username,omitempty"`
	Tags      interface{} `json:"tags,omitempty"`
	CreatedAt int         `json:"created_at,omitempty"`
}
type Interface interface {
	CreateConsumer(username string) (*Consumer, error)
	CreateConsumerKey(consumerId string) (*ConsumerKey, error)
	GetConsumerKey(consumerId, keyId string) (*ConsumerKey, error)
	GetAllConsumerKeys(consumerId string) (*Keys, error)
	RetrieveConsumerKey(consumerId string) (*ConsumerKey, error)
	DeleteConsumerKey(consumerId, consumerKeyId string) error
	DeleteConsumer(consumerId string) error
}
