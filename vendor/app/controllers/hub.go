package controllers

import (
	m "app/models"
	"log"
)

var hub *Hub

// Hub - struct for maintain clients and messages
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *SendMessage
}

// NewHub ...
func NewHub() {
	hub = &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *SendMessage),
	}
	go hub.run()
}

// infinite loop for broadcast messages
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			changeMessage(message)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// GetHub ...
func GetHub() *Hub {
	return hub
}

// changeMessage - sync message with database
func changeMessage(sm *SendMessage) {
	log.Println(sm, "changeMessage")
	message := &m.Message{ID: sm.ID, UserID: uint(sm.UserID), Message: sm.Message}
	if sm.IsDelete {
		message.DeleteMessage()
		return
	}
	messageID, err := message.SaveMessage()
	if err != nil {
		log.Println(err)
		return
	}
	message.ID = messageID
	user, err := m.GetUserByID(message.UserID)
	if err != nil {
		log.Println(err)
		return
	}
	sm.Username = user.Login
}
