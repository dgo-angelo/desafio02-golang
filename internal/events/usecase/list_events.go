package usecase

import (
	"github.com/dgo-angelo/desafio02-golang/internal/events/domain"
)

type ListEventsOutputDTO struct {
	Events []EventDTO `json:"events"`
}

type EventDTO struct {
	ID           int  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Price        float64 `json:"price"`
	ImageURL     string  `json:"image_url"`
}

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}


func (uc *ListEventsUseCase) Execute() (*ListEventsOutputDTO, error) {
	events, err := uc.repo.ListEvents()
	if err != nil {
		return nil, err
	}

	eventDTOs := make([]EventDTO, len(events))
	for i, event := range events {
		eventDTOs[i] = EventDTO{
			ID:           event.ID,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date,
			Price:        event.Price,
		}
	}

	return &ListEventsOutputDTO{Events: eventDTOs}, nil
}
