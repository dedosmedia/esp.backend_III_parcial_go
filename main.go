package main

import (
	"fmt"
	"parcial/tickets"
)

func main() {

	tickets.LoadFile("tickets.csv")
	fmt.Println("Total de registros correctos en el archivo: ", len(tickets.GetTicketList()))

	ReportTicketsByDestination("Canada")
	ReportTicketsByDestination("Portugal")
	ReportTicketsByDestination("Colombia")
	ReportTicketsByDestination("Mexico")
	ReportTicketsByDestination("No existe")

	ReportTicketsByPeriod("madrugada")
	ReportTicketsByPeriod("mañana")
	ReportTicketsByPeriod("tarde")
	ReportTicketsByPeriod("noche")
	ReportTicketsByPeriod("No existe")

	ReportTicketsPriceByDestination("Canada")
	ReportTicketsPriceByDestination("Portugal")
	ReportTicketsPriceByDestination("Colombia")
	ReportTicketsPriceByDestination("Mexico")
	ReportTicketsPriceByDestination("No existe")

}

func ReportTicketsByDestination(destination string) {
	data, err := tickets.GetTotalTickets(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Cantidad de tickets vendidos hoy a %s fue de %d\n", destination, data)
	}
}

func ReportTicketsByPeriod(period string) {
	data, err := tickets.GetCountByPeriod(period)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Cantidad de tickets vendidos hoy en el periodo %s fue de %d\n", period, data)
	}
}

func ReportTicketsPriceByDestination(destination string) {
	data, err := tickets.AverageDestination(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Promedio de precio de los tickets vendidos a %s fue de %d\n", destination, data)
	}
}
