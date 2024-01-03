package models

import "time"

type Event struct {
	Id          uint
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      uint
}

var events = []Event{}

func (e *Event) Save() {
	// todo: store in db
	events = append(events, *e)
}

func GetAllEvents() []Event {
	// todo: retrieve from db
	return events
}
