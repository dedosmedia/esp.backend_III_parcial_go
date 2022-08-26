package main

import (
	"fmt"
	"parcial/tickets"
)

func main() {

	tickets.LoadFile("tickets.csv")
	fmt.Println("Total de registros correctos en el archivo: ", len(tickets.GetTicketList()))

	fmt.Println("Lista de paises de destino")
	destinations := tickets.GetDestionations()
	for _, destination := range destinations {
		fmt.Println("\nPAIS: ", destination)
		ReportTicketsByDestination(destination)
		ReportTicketsPriceByDestination(destination)
		fmt.Println("================")
		fmt.Println()
	}

	fmt.Println("Reporte de tickets vendidos por periodo de tiempo")
	ReportTicketsByPeriod("madrugada")
	ReportTicketsByPeriod("mañana")
	ReportTicketsByPeriod("tarde")
	ReportTicketsByPeriod("noche")
	ReportTicketsByPeriod("No existe")

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
