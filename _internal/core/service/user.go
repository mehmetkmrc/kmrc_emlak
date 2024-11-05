package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/core/domain/aggregate"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/core/domain/entity"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/core/port/auth"
	"github.com/mehmetkmrc/kmrc_emlak/_internal/core/port/user"
	"strings"
)

var (
	_              user.UserServicePort = (*UserService)(nil)
	UserServiceSet                      = wire.NewSet(NewUserService)
)

type UserService struct {
	userRepo user.UserRepositoryPort
	token    auth.TokenMaker
}

func NewUserService(userRepo user.UserRepositoryPort, token auth.TokenMaker) user.UserServicePort {
	return &UserService{
		userRepo,
		token,
	}
}

func (us *UserService) Login(ctx context.Context, email, password string) (*aggregate.UserAccess, error) {
	userPassword, err := us.userRepo.GetUserPassword(ctx, email)
	if err != nil {
		return nil, err
	}
	if strings.Compare(password, userPassword) != 0 {
		return nil, errors.New("password not match")
	}
	userModel, err := us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	accessToken, publicKey, accessPayload, err := us.token.CreateToken(userModel.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshPublicKey, err := us.token.CreateRefreshToken(accessPayload)
	if err != nil {
		return nil, err
	}

	sessionModel := aggregate.NewUserAccess(userModel, accessToken, publicKey, refreshToken, refreshPublicKey)

	return sessionModel, nil
}

func (us *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	userModel, err := us.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}
