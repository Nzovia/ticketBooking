package main

import "fmt"

func main(){
	var ticket Ticket

	ticket.id =1
	ticket.number = 3
	ticket.clientName = "Nicholas"
	ticket.eventVenue = "GreenStar"

	displayTicket(ticket)

}

type Ticket struct{
	id int
	number int
	clientName string
	eventVenue string
}

func displayTicket(ticko Ticket){
	fmt.Println("Your Ticket details are.....")
	fmt.Println("Id:", ticko.id)
	fmt.Println("number:", ticko.number)
	fmt.Println("name:", ticko.clientName)
	fmt.Println("venue:",ticko.eventVenue)

}


