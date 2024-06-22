package domain

import (
	"errors"
)

var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateFuture = errors.New("event date must be in the future")
	ErrEventCapacityZero = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero = errors.New("event pice must be greater than zero")
	ErrInvalidEvent    = errors.New("invalid event data")
	ErrEventFull       = errors.New("event is full")
	ErrTicketNotFound  = errors.New("ticket not found")
	ErrTicketNotEnough = errors.New("not enough tickets available")
	ErrEventNotFound   = errors.New("event not found")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           int `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Organization string `json:"organization"`
	Rating       string `json:"rating"`
	Date        string `json:"date"`
	ImageURL     string `json:"image_url"`
	Price        float64 `json:"price"`
	Spots        []Spot `json:"spots"`
}

func(e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	

	if e.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil
}