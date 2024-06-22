package usecase

import "github.com/dgo-angelo/desafio02-golang/internal/events/domain"

type SpotDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	EventID  int    `json:"event_id"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
	TicketID string `json:"ticket_id"`
}

type TicketDTO struct {
	ID         int     `json:"id"`
	SpotID     int     `json:"spot_id"`
	TicketType string  `json:"ticket_type"`
	Price      float64 `json:"price"`
}

type Events struct {
	Events []domain.Event `json:"events"`
	Spots  []domain.Spot  `json:"spots"`
}