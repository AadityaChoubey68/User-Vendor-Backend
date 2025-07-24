package service

import (
	"context"
	"time"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
	"github.com/AadityaChoubey68/user-vendor-dashboard/storage"
)

type VendorService struct {
	storage storage.VendorStoreInterface
}

func NewVendorService(storage storage.VendorStoreInterface) *VendorService {
	return &VendorService{
		storage: storage,
	}
}

func (s *VendorService) CreateVendor(ctx context.Context, req *models.VendorSignUpRequest) (*models.VendorSignUpRequest, error) {
	Vendor := models.Vendor{
		Name:      req.Name,
		Emial:     req.Emial,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	createdVendor, err := s.storage.CreateVendor(ctx, Vendor)
	if err != nil {
		return nil, err
	}
	return &createdVendor, err
}

func (s *VendorService) LoginVendor(ctx context.Context, req models.VendorLoginRequest) (models.VendorLoginResponse, error) {
	Vendor, err := s.storage.GetVendorByEmail(ctx, req.Email)

	if Vendor.Password != req.Password {
		return models.VendorLoginResponse{}, err
	}

	return models.VendorLoginResponse{
		Name:  Vendor.Name,
		Email: Vendor.Emial,
	}, nil
}
