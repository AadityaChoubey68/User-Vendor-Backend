package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
)

type VdStore struct {
	db *sql.DB
}

func VdNew(db *sql.DB) *VdStore {
	return &VdStore{db: db}
}

func (s *VdStore) CreateVendor(ctx context.Context, user models.Vendor) (models.VendorSignUpRequest, error) {
	var returnedVendor models.VendorSignUpRequest
	query := `INSERT INTO vendors (name, email, password, created_at)
	          VALUES ($1, $2, $3, $4)
			  RETURNING name, email, password`
	err := s.db.QueryRowContext(ctx, query,
		user.Name,
		user.Emial,
		user.Password,
		user.CreatedAt,
	).Scan(
		&returnedVendor.Name,
		&returnedVendor.Emial,
		&returnedVendor.Password,
	)
	if err != nil {
		log.Println("Error in user storage layer")
		return returnedVendor, err
	}
	return returnedVendor, err
}

func (s *VdStore) GetVendorByEmail(ctx context.Context, email string) (models.Vendor, error) {
	query := `SELECT id,name,email,password,created_at FROM vendors WHERE email=$1`
	var Vendor models.Vendor
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&Vendor.Id,
		&Vendor.Name,
		&Vendor.Emial,
		&Vendor.Password,
		&Vendor.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Vendor, fmt.Errorf("user not found")
		}
		return Vendor, err
	}
	return Vendor, nil
}
