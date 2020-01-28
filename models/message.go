package models

import "fmt"

// Message jest standardowym modelem
// dla krotkich wiadomosci dla informacji do uzytkownika
// zazwyczaj ma kod 200
type Message struct {
	Message string `json:"message"`
}

// NewMessage to konstruktor dla Message
func NewMessage(message string) *Message {
	msg := new(Message)
	msg.Message = message
	return msg
}

// Msg200 zwraca info dla gin'a
func Msg200(message string) (int, *Message) {
	msg := NewMessage(message)
	return 200, msg
}

// Msg200f to wrapper dla Msg200 tylko ze z "f"
func Msg200f(message string, a ...interface{}) (int, *Message) {
	info := fmt.Sprintf(message, a...)
	msg := NewMessage(info)
	return 200, msg
}
