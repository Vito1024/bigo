package protocol

type Request struct {
	Header header `json:"header"`
	Body   []byte `json:"body"`
}

type header struct {
	ContentLength int `json:"content_length"`
}
