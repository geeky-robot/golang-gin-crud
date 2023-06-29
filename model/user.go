package model

import "github.com/google/uuid"

type User struct {
	Id      uuid.UUID
	Name    string
	Dob     string
	Age     int64
	Address Address
}
