package controllers

import (
	m "app/models"
	"bytes"
	"log"
	"time"

	soc "github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client - struct for chat
type Client struct {
	m.User
	con  *soc.Conn
	send chan []byte
}

func (c *Client) read() {
	hub := GetHub()
	defer func() {
		hub.unregister <- c
		c.con.Close()
	}()
	c.con.SetReadLimit(maxMessageSize)
	c.con.SetReadDeadline(time.Now().Add(pongWait))
	c.con.SetPongHandler(func(string) error { c.con.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.con.ReadMessage()
		if err != nil {
			if soc.IsUnexpectedCloseError(err, soc.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		hub.broadcast <- message
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.con.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.con.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.con.WriteMessage(soc.CloseMessage, []byte{})
				return
			}

			w, err := c.con.NextWriter(soc.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.con.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.con.WriteMessage(soc.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
