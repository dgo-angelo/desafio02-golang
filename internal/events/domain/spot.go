package domain

import (
	"errors"
)

type SpotStatus string

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotNameTwoCharacters   = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameRequired        = errors.New("spot name is required")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndWithLetter   = errors.New("spot name must start with a number")
)

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       int `json:"id"`
	EventID  int  `json:"event_id"`
	Name     string `json:"name"`
	Status   SpotStatus `json:"status"`
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}
	// Validate if the spot name is in the correct format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithLetter
	}
	return nil
}