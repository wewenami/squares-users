package models

import "time"

type User struct {
	Id         string
	FirstName  string
	LastName   string
	Patronymic string
	Phone      string
	Email      string
	IsEnabled  bool
	PhotoURL   string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}
