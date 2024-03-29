package models

import (
	"time"

	"github.com/criticalsession/go-rest/db"
)

type Event struct {
	Id            uint
	Name          string    `binding:"required"`
	Description   string    `binding:"required"`
	Location      string    `binding:"required"`
	DateTime      time.Time `binding:"required"`
	UserId        uint
	Registrations []Registration
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, date_time, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	e.Id = uint(id)
	events = append(events, *e)

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, date_time, user_id FROM events"
	res, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	events := []Event{}
	for res.Next() {
		e := Event{}
		err := res.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}

		e.Registrations, err = e.GetRegistrations()
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id uint) (*Event, error) {
	query := "SELECT id, name, description, location, date_time, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		return nil, err
	}

	e.Registrations, err = e.GetRegistrations()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date_time = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)
	return err
}

func (e *Event) Register(userId uint) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)
	return err
}

func (e *Event) Unregister(userId uint) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	_, err := db.DB.Exec(query, e.Id, userId)
	return err
}

func (e *Event) GetRegistrations() ([]Registration, error) {
	query := "SELECT id,event_id, user_id FROM registrations WHERE event_id = ?"
	registrations := []Registration{}
	rows, err := db.DB.Query(query, e.Id)
	if err != nil {
		return []Registration{}, err
	}

	for rows.Next() {
		r := Registration{}
		err := rows.Scan(&r.Id, &r.EventId, &r.UserId)
		if err != nil {
			return []Registration{}, err
		}

		registrations = append(registrations, r)
	}

	return registrations, nil
}
