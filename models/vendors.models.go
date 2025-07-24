package models

import "time"

type Vendor struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Emial     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type VendorSignUpRequest struct {
	Name     string `json:"name"`
	Emial    string `json:"email"`
	Password string `json:"password"`
}

type VendorLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VendorLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
