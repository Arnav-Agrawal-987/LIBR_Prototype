package model

import "github.com/google/uuid"

type Message struct {
	ID        uuid.UUID `json:"uuid"`
	Content   string    `json:"content"`
	Timestamp int       `json:"timestamp"`
	Status    string    `json:"status"`
}

type ModResponse struct {
	ModID        int
	MessageID    uuid.UUID
	Status       int
	ResponseTime int
}

type ApiResponse struct {
	ID        uuid.UUID `json:"uuid"`
	Timestamp int       `json:"timestamp"`
	Status    string    `json:"status"`
}
