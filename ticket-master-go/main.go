package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"ticketmaster/engine"
)

type Ticket struct {
	EventID    string `json:"event_id"`
	EventName  string `json:"event_name"`
	TicketLink string `json:"ticket_link"`
	Image      string `json:"image"`
	Country    string `json:"country"`
}

type TicketData struct {
	ES []Ticket `json:"es"`
	UK []Ticket `json:"uk"`
}

func main() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var data TicketData

	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	ticketsSpain := data.ES
	ticketsUk := data.UK

	for _, ticket := range ticketsSpain {
		wg.Add(1)
		go engine.MasterSpain(ticket.EventID, ticket.EventName, ticket.TicketLink, ticket.Image, ticket.Country)
	}

	for _, ticket := range ticketsUk {
		wg.Add(1)
		go engine.MasterUk(ticket.EventID, ticket.EventName, ticket.TicketLink, ticket.Image, ticket.Country)
	}

	wg.Wait()
}

//check for memory leaks
// // pid := "36051"
// // https://www.ticketmaster.es/event/christian-nodal-entradas/36051
