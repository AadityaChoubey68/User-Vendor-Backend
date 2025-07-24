package service

import (
	"context"
	"time"

	"github.com/AadityaChoubey68/user-vendor-dashboard/models"
	"github.com/AadityaChoubey68/user-vendor-dashboard/storage"
)

type Service struct {
	storage storage.UserStorageInterface
}

func NewUserService(storage storage.UserStorageInterface) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateUser(ctx context.Context, req *models.UserSignUpRequest) (*models.UserSignUpRequest, error) {
	user := models.User{
		Name:      req.Name,
		Emial:     req.Emial,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	createdUser, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &createdUser, err
}

func (s *Service) LoginUser(ctx context.Context, req models.UserLoginRequest) (models.UserLoginResponse, error) {
	user, err := s.storage.GetUserByEmail(ctx, req.Email)

	if user.Password != req.Password {
		return models.UserLoginResponse{}, err
	}

	return models.UserLoginResponse{
		Name:  user.Name,
		Email: user.Emial,
	}, nil
}
