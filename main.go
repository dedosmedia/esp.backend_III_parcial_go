package main

import (
	"fmt"
	"log"
	"os"
	"parcial/tickets"
	"strings"
)

func main() {

	// Leer el fichero desde disco
	data, err := os.ReadFile("tickets.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Separar cada línea del fichero
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		// recorrer cada línea y convertirla en Ticket
		_, err := tickets.ConvertLineToTicket(line)
		if err != nil {
			fmt.Printf("No se procesará la siguiente línea: ´%s´ porque %s\n", line, err)
			continue
		}
	}
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
