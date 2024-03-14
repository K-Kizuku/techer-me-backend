package user

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/user"
)

type IUserService interface {
	CreateUserByFirebaseID(ctx context.Context, firebaseID string) error
	CreateUserDetailByFirebaseID(ctx context.Context, firebaseID string, name string, imageURL string) error
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

func (s *Service) CreateUserDetailByFirebaseID(ctx context.Context, firebaseID string, name string, imageURL string) error {
	user := &entity.User{
		ID:       firebaseID,
		Name:     name,
		ImageURL: imageURL,
	}
	if err := s.userRepo.CreateDetail(ctx, user); err != nil {
		return err
	}
	return nil
}
