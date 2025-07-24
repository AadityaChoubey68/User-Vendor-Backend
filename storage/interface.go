package storage

import (
	"context"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
)

type UserStorageInterface interface {
	CreateUser(ctx context.Context, user models.User) (models.UserSignUpRequest, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type VendorStoreInterface interface {
	CreateVendor(ctx context.Context, user models.Vendor) (models.VendorSignUpRequest, error)
	GetVendorByEmail(ctx context.Context, email string) (models.Vendor, error)
}
