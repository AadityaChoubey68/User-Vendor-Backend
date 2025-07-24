package service

import (
	"context"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, req *models.UserSignUpRequest) (*models.UserSignUpRequest, error)
	LoginUser(ctx context.Context, req models.UserLoginRequest) (models.UserLoginResponse, error)
}

type VendorServiceInterface interface {
	CreateVendor(ctx context.Context, req *models.VendorSignUpRequest) (*models.VendorSignUpRequest, error)
	LoginVendor(ctx context.Context, req models.VendorLoginRequest) (models.VendorLoginResponse, error)
}
