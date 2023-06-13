package models

import (
	"database/sql"
	"errors"
	"time"
)

type Ticket struct {
	ID          int
	Title       string
	Description string
	Status      int
	Created     time.Time
	Resolved    time.Time
}

type TicketModel struct {
	DB *sql.DB
}

func (m *TicketModel) ViewTicket(id int) (*Ticket, error) {
	statement := `SELECT ID, Title, Description, Status, Created, Resolved FROM
		tickets WHERE ID = $1`

	row := m.DB.QueryRow(statement, id)

	t := &Ticket{}

	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Created, &t.Resolved)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return t, nil
}
