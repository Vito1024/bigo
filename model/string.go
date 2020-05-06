package model

type String struct {
	*baseType
}

func NewString() *String {
	return &String{newbaseType()}
}

// format of string GET command
type GETFormat struct {
	Key string `json:"key"`
}

// format of string SET command
type SETFormat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
