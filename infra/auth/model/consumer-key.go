package model

type ConsumerKey struct {
	Tags      interface{} `json:"tags"`
	CreatedAt int         `json:"created_at"`
	Key       string      `json:"key"`
	Id        string      `json:"id"`
	Ttl       interface{} `json:"ttl"`
	Consumer  struct {
		Id string `json:"id"`
	} `json:"consumer"`
}

type Keys struct {
	Data []ConsumerKey `json:"data"`
	Next interface{}   `json:"next"`
}
