package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(ctx context.Context, user models.User) (models.UserSignUpRequest, error) {
	var returnedUser models.UserSignUpRequest
	query := `INSERT INTO users (name, email, password, created_at)
	          VALUES ($1, $2, $3, $4)
			  RETURNING name, email, password`
	err := s.db.QueryRowContext(ctx, query,
		user.Name,
		user.Emial,
		user.Password,
		user.CreatedAt,
	).Scan(
		&returnedUser.Name,
		&returnedUser.Emial,
		&returnedUser.Password,
	)
	if err != nil {
		log.Println("Error in user storage layer")
		return returnedUser, err
	}
	return returnedUser, err
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT id,name,email,password,created_at FROM users WHERE email=$1`
	var user models.User
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Emial,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}
	return user, nil
}
