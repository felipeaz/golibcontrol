package model

type UserSession struct {
	UserId string `json:"userId" binding:"required"`
	KeyId  string `json:"keyId" binding:"required"`
}
