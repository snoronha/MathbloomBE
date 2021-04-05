package models

type Ticket struct {
	TicketType string `json:"ticketType" sql:"type:varchar(16);unique_index"`
	TicketId   uint   `json:"ticketId"`
}
