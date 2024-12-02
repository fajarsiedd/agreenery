package auth

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/repositories/auth"
)

type authService struct {
	repository       auth.AuthRepository
	jwtConfig        *middlewares.JWTConfig
	jwtRefreshConfig *middlewares.JWTConfig
}

func NewAuthService(r auth.AuthRepository, jwtConfig, jwtRefreshConfig *middlewares.JWTConfig) *authService {
	return &authService{
		repository:       r,
		jwtConfig:        jwtConfig,
		jwtRefreshConfig: jwtRefreshConfig,
	}
}

func (service authService) Login(user entities.User) (entities.User, error) {
	result, err := service.repository.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	match, err := helpers.CompareStringWithEncodedHash(user.Credential.Password, result.Credential.Password)

	isFailed := err != nil || !match
	if isFailed {
		return entities.User{}, constants.ErrIncorrectPassword
	}

	accessToken, err := service.jwtConfig.GenerateToken(result.ID, string(result.Credential.Role))
	if err != nil {
		return entities.User{}, err
	}

	refreshToken, err := service.jwtRefreshConfig.GenerateRefreshToken(result.ID, string(result.Credential.Role))
	if err != nil {
		return entities.User{}, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}

func (service authService) Register(user entities.User) (entities.User, error) {
	config := &helpers.ArgonConfig{
		Memory:     64 * 1024,
		Iterations: 3,
		Pararelism: 2,
		SaltLength: 16,
		KeyLength:  32,
	}

	var err error
	user.Credential.Password, err = helpers.HashString(user.Credential.Password, config)
	if err != nil {
		return entities.User{}, err
	}

	result, err := service.repository.Register(user)
	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}

func (service authService) GetNewTokens(claims *middlewares.JWTCustomClaims, refreshToken string) (entities.User, error) {
	newAccessToken, err := service.jwtConfig.GenerateToken(claims.UserID, claims.Role)
	if err != nil {
		return entities.User{}, err
	}

	newRefreshToken, err := service.jwtRefreshConfig.GenerateRefreshToken(claims.UserID, claims.Role)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (service authService) GetProfile(id string) (entities.User, error) {
	return service.repository.FindUser(id)
}

func (service authService) UpdateProfile(user entities.User, selectedFields []string) (entities.User, error) {
	return service.repository.UpdateUser(user, selectedFields)
}
