package models

import (
	"time"

)

type Ticket struct{
	id int `gorm:"primaryKey" json:"id"`
	ticketTitle string `json:"ticketTitle"`
	ticketDescription string `json:"ticketDescription"`
	ticketPrice int `json:"ticketPrice"`
	eventDate time.Time `json:"eventdate"`

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