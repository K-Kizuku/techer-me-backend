package user

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/user"
)

type IUserService interface {
	CreateUserByFirebaseID(ctx context.Context, firebaseID string) error
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
