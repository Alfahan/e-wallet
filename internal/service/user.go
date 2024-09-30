package service

import (
	"context"
	"e-wallet/domain"
	"e-wallet/domain/dto"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Authenticate implements domain.UserService.
func (u *userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

}

// ValidateToken implements domain.UserService.
func (u *userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	panic("unimplemented")
}
