package main

import "fmt"

type BusTicket struct {
	Passanger     string
	DepartureCity string
	ArrivalCity   string
	DepartureTime string
	BusType       string
	NationalCode  string
	Price         string
}

type PlaneTicket struct {
	Passanger        string
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	PlaneType        string
	PassportID       string
	Price            string
}

type TicketPrinter interface {
	PrintTicket()
}

func (bus BusTicket) PrintTicket() {
	fmt.Printf("Bus Ticket:\n Price : %s\n", bus.Price)
}

func (plane PlaneTicket) PrintTicket() {
	fmt.Printf("Plane Ticket:\n Price : %s\n", plane.Price)
}

func main() {

	var tickets []TicketPrinter
	bus := BusTicket{
		"Hossein", "Isfahan", "Tehran", "20:55", "VIP", "1254789657", "350000",
	}
	plane := PlaneTicket{
		"Hossein", "London", "Tehran", "10:30", "21:20", "B-Class", "5478J658Q", "1200000",
	}

	tickets = append(tickets, bus, plane)
	for _, value := range tickets {
		value.PrintTicket()
	}
}
