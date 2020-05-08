package model

type BigoRequest struct {
	CommandName string     `json:"name"`
	Args        []string   `json:"args"`
	ClientInfo  ClientInfo `json:"client_info"`
}

type ClientInfo struct {
	ClientId string `json:"client_id"`
}
