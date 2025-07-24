package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Emial     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserSignUpRequest struct {
	Name     string `json:"name"`
	Emial    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
