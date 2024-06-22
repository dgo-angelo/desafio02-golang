package usecase

import (
	"github.com/dgo-angelo/desafio02-golang/internal/events/domain"
)

type BuyTicketsInputDTO struct {
	Spots      []string `json:"spots"`
}

type BuyTicketsOutputDTO struct {
	Tickets []TicketDTO `json:"tickets"`
}


type BuyTicketsUseCase struct {
	repo           domain.EventRepository
}

func NewBuyTicketsUseCase(repo domain.EventRepository) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		repo:           repo,
		
	}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO, eventID int) (*BuyTicketsOutputDTO, error) {
	// Verifica o evento
	_, err := uc.repo.FindEventByID(eventID)
	if err != nil {
		return nil, err
	}
	// Salva os ingressos no banco de dados
	for _, spot := range input.Spots {
		err = uc.repo.ReserveSpot(spot, eventID)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
