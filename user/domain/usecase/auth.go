package usecase

import (
	"context"
	"errors"
	"event_sourcing_user/application/command"
	"event_sourcing_user/application/writeOutput"
	"event_sourcing_user/constant"
	"event_sourcing_user/domain/value_object"
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/package/encrypt_password"
	"event_sourcing_user/package/ierror"
	"event_sourcing_user/package/jwt_utils"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/proto/user"

	"go.uber.org/zap"
)

type AuthUsecase interface {
	Login(ctx context.Context, command *command.LoginCommand) (*writeOutput.LoginOutput, error)
	Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*writeOutput.LoginOutput, error)
}

type authUsecase struct {
	userRepository  repository.IUserRepository
	jwtUtils        *jwt_utils.JWTUtils
	encryptPassword *encrypt_password.ArgonParam
}

func NewAuthUsecase(config *constant.Config, factoryRepository repository.IRepositoryFactory) (AuthUsecase, error) {
	encryptPassword, err := encrypt_password.NewArgonParam()
	jwtUtils := jwt_utils.NewJWTUtils(
		config.Security.JWTAccessSecret,
		config.Security.JWTRefreshSecret,
		int32(config.Security.JWTAccessExpiration),
		int32(config.Security.JWTRefreshExpiration),
	)
	if err != nil {
		return nil, ierror.Error(err)
	}
	return &authUsecase{
		userRepository:  factoryRepository.UserRepository(),
		encryptPassword: encryptPassword,
		jwtUtils:        jwtUtils,
	}, nil
}

func (u *authUsecase) Login(ctx context.Context, command *command.LoginCommand) (*writeOutput.LoginOutput, error) {
	log := logger.FromContext(ctx)
	user, err := u.userRepository.GetUserByEmail(command.Email)
	if err != nil {
		log.Error("Error getting user by email", zap.Error(err), zap.String("email", command.Email))
		return nil, ierror.Error(err)
	}

	password, err := value_object.NewAndValidatePassword(command.Password)
	if err != nil {
		log.Error("Invalid password", zap.Error(err))
		return nil, ierror.Error(err)
	}

	verify, err := u.encryptPassword.VerifyPassword(user.Password(), password.Value())
	if err != nil {
		log.Error("Error verifying password", zap.Error(err))
		return nil, ierror.Error(err)
	}
	if !verify {
		return nil, ierror.Error(errors.New("invalid password"))
	}

	accessToken, refreshToken, _, _, err := u.jwtUtils.GenerateToken(user)
	if err != nil {
		log.Error("Error generating token", zap.Error(err))
		return nil, ierror.Error(err)
	}

	return &writeOutput.LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *authUsecase) Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*writeOutput.LoginOutput, error) {
	return nil, nil
}
