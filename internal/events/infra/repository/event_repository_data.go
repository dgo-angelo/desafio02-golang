package repository

import (
	"log"
	"slices"
	"strconv"

	"github.com/dgo-angelo/desafio02-golang/internal/events/domain"
)

type Events struct {
	Events []domain.Event `json:"events"`
	Spots  []domain.Spot         `json:"spots"`
}

type Spot struct {
	ID      int    `json:"id"`
	EventID int    `json:"event_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

// NewMysqlEventRepository creates a new MySQL event repository.
func NewDataEventRepository(db *Events) (domain.EventRepository, error) {
	return &dataEventRepository{db: db}, nil
}

// Definição da estrutura dataEventRepository
type dataEventRepository struct {
	db *Events
}

// FindSpotByName implements domain.EventRepository.
func (r *dataEventRepository) FindSpotByName(eventID int, spotName string) (*domain.Spot, error) {
	panic("unimplemented")
}

// FindSpotsByEventID implements domain.EventRepository.
func (r *dataEventRepository) FindSpotsByEventID(eventID int) ([]domain.Spot, error) {
	spots := r.db.Spots
	filtered := []domain.Spot{}
	for _, spot := range spots {
		if spot.EventID == eventID {
			filtered = append(filtered, spot)
		}
	}
	return filtered, nil;
}


// ListEvents returns all events with their associated spots and tickets.
func (r *dataEventRepository) ListEvents() ([]domain.Event, error) {
	eventMap := make(map[string]domain.Event)
	for _, event := range r.db.Events {
		var eventId = strconv.Itoa(event.ID)
		eventMap[eventId] = event
	}

	return r.db.Events, nil
}

// FindEventByID returns an event by its ID, including associated spots and tickets.
func (r *dataEventRepository) FindEventByID(eventID int) (domain.Event, error) {
	var data = r.db.Events
	index := slices.IndexFunc(r.db.Events, func(e domain.Event) bool { return e.ID == eventID })
	return data[index], nil
}

// ReserveSpot implements domain.EventRepository.
func (r *dataEventRepository) ReserveSpot(spotName string, eventID int) error {
	
	index := slices.IndexFunc(r.db.Spots, func(s domain.Spot) bool { return s.EventID == eventID && s.Name == spotName })

	if r.db.Spots[index].Status != "reserved" {
		r.db.Spots[index].Status = "reserved"
	}else{
		return domain.ErrSpotAlreadyReserved
	}

	log.Println(spotName)
	return nil
}