package model

import "time"

type User struct {
	ID        int
	Nama      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
