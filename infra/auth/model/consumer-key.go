package model

type CreateKeyBody struct {
	Secret string `json:"secret"`
}

type ConsumerKey struct {
	Key       string `json:"key"`
	CreatedAt int    `json:"created_at"`
	Id        string `json:"id"`
	Consumer  struct {
		Id string `json:"id"`
	} `json:"consumer"`
	Tags         interface{} `json:"tags"`
	RsaPublicKey interface{} `json:"rsa_public_key"`
	Algorithm    string      `json:"algorithm"`
	Secret       string      `json:"secret"`
}

type Keys struct {
	Data []ConsumerKey `json:"data"`
	Next interface{}   `json:"next"`
}
