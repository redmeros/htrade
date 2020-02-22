package models

import "encoding/json"

type WsMessage struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (m *WsMessage) Json() []byte {
	js, _ := json.Marshal(&m)
	return js
}

func NewNotify(msg string) *WsMessage {
	m := WsMessage{
		Action:  "notify",
		Data:    nil,
		Message: msg,
	}
	return &m
}

func NewError(msg string, data interface{}) *WsMessage {
	m := WsMessage{
		Action:  "error",
		Data:    data,
		Message: msg,
	}
	return &m
}
