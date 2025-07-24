package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
	"github.com/AadityaChoubey68/user-vendor-dashboard/service"
)

type Userhandler struct {
	service service.UserServiceInterface
}

func NewHandler(service service.UserServiceInterface) *Userhandler {
	return &Userhandler{
		service: service,
	}
}

func (h *Userhandler) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserSignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.CreateUser(context.Background(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *Userhandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var requestedUser models.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestedUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	user, err := h.service.LoginUser(ctx, requestedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
