package repository

import (
	m "github.com/geeky-robot/golang-gin-crud/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	CreateUser(user m.User) m.User
	CreateUsers(users []m.User) []m.User
	UpdateUser(user m.User) m.User
	GetUser(id uuid.UUID) m.User
	GetUsers() []m.User
	DeleteUser(id uuid.UUID) bool
}
