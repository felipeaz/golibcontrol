package model

type UserSession struct {
	ConsumerId    string `json:"consumerId" binding:"required"`
	ConsumerKeyId string `json:"consumerKeyId"`
}
