package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/dgo-angelo/desafio02-golang/internal/events/infra/http"
	"github.com/dgo-angelo/desafio02-golang/internal/events/infra/repository"
	"github.com/dgo-angelo/desafio02-golang/internal/events/usecase"
)



var database *repository.Events

// @title Events API
// @version 1.0
// @description This is a server for managing events. Imersão Full Cycle
// @host localhost:8080
// @BasePath /
func main() {
	// Configuração do banco de dados

	jsonData, err := os.Open("./data.json")
 	byteValue, _ := io.ReadAll(jsonData)

 	err = json.Unmarshal(byteValue, &database)
    if err != nil {
        fmt.Println(err)
		log.Fatal(err)
    }
    defer jsonData.Close()

	if err != nil {
        fmt.Println(err)
    }

	// Repositório
	eventRepo, err := repository.NewDataEventRepository(database)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)

	// Handlers HTTP
	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()

	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /event/{eventID}/reserve", eventsHandler.BuyTickets)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Recebido sinal de interrupção, iniciando o graceful shutdown
		log.Println("Recebido sinal de interrupção, iniciando o graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro no graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Iniciando o servidor HTTP
	log.Println("Servidor HTTP rodando na porta 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor HTTP finalizado")
}
