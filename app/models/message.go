package models

// Message describe a message sent and receive between clients
type Message struct {
	ID        string `json:"ID"`
	Sender    string
	Timestamp int64
	Content   string `json:"content"`
}

// RawClientMessage describe a message sent from a client
type RawClientMessage struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
