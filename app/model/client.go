package model

import (
	"github.com/gorilla/websocket"
)

// Client represents a single chatting user.
type Client struct {
	ID string

	// socket is the web socket for this client.
	Socket *websocket.Conn

	// send is a channel on which messages are sent.
	Send chan []byte

	// room is the room this client is chatting in.
	Room *Room
}

func (c *Client) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		c.Room.Forward <- &RawClientMessage{Sender: c.ID, Content: string(msg)}
	}
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for msg := range c.Send {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
