package models

import "time"

type Terminal struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TerminalRequest struct {
	Name string `json:"name"`
}
