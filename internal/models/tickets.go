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

func (m *TicketModel) Latest() ([]*Ticket, error) {

	stmt := `SELECT id, Title, Description, Created, Status FROM tickets
    WHERE Status = 1 ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt) // execute statement on database, saves them to the rows variable
	if err != nil {
		return nil, err
	}

	defer rows.Close() // ensure that the connection pool is closed at the end of runtime

	Tickets := []*Ticket{} // create slice of Tickets to be filled by those acquired by statement

	for rows.Next() { // while there still is a row to be seen

		s := &Ticket{} // points to a new Ticket struct

		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.Created, &s.Status) // pull Content from each Ticket in row
		if err != nil {
			return nil, err
		}

		Tickets = append(Tickets, s) // add each Ticket (if no error) to the slice
	}

	if err = rows.Err(); err != nil { // if an error was encountered while looping
		return nil, err
	}

	return Tickets, nil // finally, return the slice of Tickets + no error
}
