package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID int) (Event, error)
	FindSpotsByEventID(eventID int) ([]Spot, error)
	FindSpotByName(eventID int, spotName string) (*Spot, error)
	ReserveSpot(spotName string, eventID int) error
}