package usecase

import (
	"github.com/dgo-angelo/desafio02-golang/internal/events/domain"
)

type ListSpotsInputDTO struct {
	EventID int `json:"event_id"`
}

type ListSpotsOutputDTO struct {
	Event EventDTO  `json:"event"`
	Spots []SpotDTO `json:"spots"`
}

type ListSpotsUseCase struct {
	repo domain.EventRepository
}

func NewListSpotsUseCase(repo domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repo: repo}
}

func (uc *ListSpotsUseCase) Execute(input ListSpotsInputDTO) (*ListSpotsOutputDTO, error) {
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots, err := uc.repo.FindSpotsByEventID(input.EventID)
	if err != nil {
		return nil, err
	}

	spotDTOs := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotDTOs[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
		}
	}

	eventDTO := EventDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date,
		Price:        event.Price,
		ImageURL:     event.ImageURL,
	}

	return &ListSpotsOutputDTO{Event: eventDTO, Spots: spotDTOs}, nil
}
