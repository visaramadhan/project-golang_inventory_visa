package model

import (
	"time"
)

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	PhoneNumber      string    `json:"phone_number"`
	RegistrationDate time.Time `json:"regist_date"`
	LastLogin        time.Time `json:"last_login"`
}
