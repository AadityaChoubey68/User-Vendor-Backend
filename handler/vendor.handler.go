package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
	"github.com/AadityaChoubey68/user-vendor-dashboard/service"
)

type Vendorhandler struct {
	service service.VendorServiceInterface
}

func NewVendorHandler(service service.VendorServiceInterface) *Vendorhandler {
	return &Vendorhandler{
		service: service,
	}
}

func (h *Vendorhandler) VendorSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req models.VendorSignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdVendor, err := h.service.CreateVendor(context.Background(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error : ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdVendor)
}

func (h *Vendorhandler) VendorLogin(w http.ResponseWriter, r *http.Request) {
	var requestedVendor models.VendorLoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestedVendor)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	Vendor, err := h.service.LoginVendor(ctx, requestedVendor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Vendor)
}
