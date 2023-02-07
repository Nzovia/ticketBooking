package models

import (
	"time"

)

type Ticket struct{
	id int `gorm:"primaryKey"`
	ticketTitle string
	ticketDescription string
	ticketPrice int
	eventDate time.Time

}

type Attendee struct{
	id int `gorm:"primaryKey"`
	firstName string
	secondName string
	mobileNumber string
	totalTickets int
	Ticket Ticket `has_many:"tickets" order_by:"ticketTitle"`

	//add one to many relationship between ticket amd the client
}