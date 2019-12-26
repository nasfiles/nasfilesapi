package api

import "time"

//User is a struct holding important details about every account
type User struct {
	UID      string
	Username string
	Email    string
	Name     string
	Created  time.Time
}

//UserService ...
type UserService interface {
	Add(u *User) error
	Delete(uid string) error
}
