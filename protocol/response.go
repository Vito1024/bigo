package protocol

type Response struct {
	Header responseHeader `json:"header"`
	Body   []byte         `json:"body"`
}

type responseHeader struct {
	Status        int  `json:"status"`
	ContentLength uint `json:"content_length"`
}
