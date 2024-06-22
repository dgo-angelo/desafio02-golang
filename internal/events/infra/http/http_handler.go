package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dgo-angelo/desafio02-golang/internal/events/usecase"
)

// EventsHandler handles HTTP requests for events.
type EventsHandler struct {
	listEventsUseCase  *usecase.ListEventsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	buyTicketsUseCase  *usecase.BuyTicketsUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
}

// NewEventsHandler creates a new EventsHandler.
func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		getEventUseCase:    getEventUseCase,
		buyTicketsUseCase:  buyTicketsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
	}
}

// ListSpots lists spots for an event.
// @Summary List spots for an event
// @Description List all spots for a specific event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {object} usecase.ListSpotsOutputDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /events/{eventID}/spots [get]
func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	eventIDString, _ := strconv.Atoi(eventID)
	input := usecase.ListSpotsInputDTO{EventID: eventIDString}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// ListEvents handles the request to list all events.
// @Summary List all events
// @Description Get all events with their details
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} usecase.ListEventsOutputDTO
// @Failure 500 {object} string
// @Router /events [get]
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetEvent handles the request to get details of a specific event.
// @Summary Get event details
// @Description Get details of an event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {object} usecase.GetEventOutputDTO
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /events/{eventID} [get]
func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	eventIDString, _ := strconv.Atoi(eventID)
	input := usecase.GetEventInputDTO{ID: eventIDString}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}


// BuyTickets handles the request to buy tickets for an event.
// @Summary Buy tickets for an event
// @Description Buy tickets for a specific event
// @Tags Events
// @Accept json
// @Produce json
// @Param input body usecase.BuyTicketsInputDTO true "Input data"
// @Success 200 {object} usecase.BuyTicketsOutputDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /checkout [post]
func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {

	var input usecase.BuyTicketsInputDTO
	eventID := r.PathValue("eventID")
	eventIDString, _ := strconv.Atoi(eventID)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(input)
	output, err := h.buyTicketsUseCase.Execute(input,eventIDString )

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// writeErrorResponse writes an error response in JSON format
func (h *EventsHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// CreateSpotsRequest represents the input for creating spots.
type CreateSpotsRequest struct {
	NumberOfSpots int `json:"number_of_spots"`
}
