package model

// Message describe a message sent and receive between clients
type Message struct {
	ID        string `json:"Id"`
	Sender    string
	Timestamp int
	Content   string `json:"content"`
}
