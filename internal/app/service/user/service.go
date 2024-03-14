package user

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/user"
)

type IUserService interface {
	CreateUserByFirebaseID(ctx context.Context, firebaseID string) error
	CreateUserDetailByFirebaseID(ctx context.Context, input schema.CreateUserInput) error
	GetByID(ctx context.Context, userID string) (*entity.User, error)
}

type Service struct {
	userRepo user.IRepository
}

func New(userRepo user.IRepository) IUserService {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) CreateUserByFirebaseID(ctx context.Context, firebaseID string) error {
	if err := s.userRepo.Create(ctx, firebaseID); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateUserDetailByFirebaseID(ctx context.Context, input schema.CreateUserInput) error {
	user := &entity.User{
		ID:       input.UserID,
		Name:     input.Name,
		ImageURL: input.ImageURL,
	}
	if err := s.userRepo.CreateDetail(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	user, err := s.userRepo.SelectByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
