package types

import "time"

type User struct {
	Id        int
	FirstName string
	Email     string
	Password  string
	CreatedAt time.Time
}

type RegisterPayload struct {
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore interface {
	CreateUser(User) error
	GetUserByGmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
}
