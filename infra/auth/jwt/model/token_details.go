package model

// TokenDetails contains JWT Tokens information
type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}
