package model

type Consumer struct {
	Id        string      `json:"id,omitempty"`
	CustomId  string      `json:"custom_id,omitempty"`
	Username  string      `json:"username,omitempty"`
	Tags      interface{} `json:"tags,omitempty"`
	CreatedAt int         `json:"created_at,omitempty"`
}
