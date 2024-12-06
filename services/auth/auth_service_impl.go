package auth

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/repositories/auth"
	"mime/multipart"
	"strings"
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

func (s authService) Login(user entities.User) (entities.User, error) {
	result, err := s.repository.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	match, err := helpers.CompareStringWithEncodedHash(user.Credential.Password, result.Credential.Password)

	isFailed := err != nil || !match
	if isFailed {
		return entities.User{}, constants.ErrIncorrectPassword
	}

	accessToken, err := s.jwtConfig.GenerateToken(result.ID, string(result.Credential.Role))
	if err != nil {
		return entities.User{}, err
	}

	refreshToken, err := s.jwtRefreshConfig.GenerateRefreshToken(result.ID, string(result.Credential.Role))
	if err != nil {
		return entities.User{}, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}

func (s authService) Register(user entities.User) (entities.User, error) {
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

	result, err := s.repository.Register(user)
	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}

func (s authService) GetNewTokens(claims *middlewares.JWTCustomClaims, refreshToken string) (entities.User, error) {
	newAccessToken, err := s.jwtConfig.GenerateToken(claims.UserID, claims.Role)
	if err != nil {
		return entities.User{}, err
	}

	newRefreshToken, err := s.jwtRefreshConfig.GenerateRefreshToken(claims.UserID, claims.Role)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s authService) GetProfile(id string) (entities.User, error) {
	return s.repository.GetUser(id)
}

func (s authService) UpdateProfile(user entities.User, selectedFields []string) (entities.User, error) {
	return s.repository.UpdateUser(user, selectedFields)
}

func (s authService) UploadProfilePhoto(file multipart.File, userID string) (entities.User, error) {
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return entities.User{}, err
	}

	var oldObj string
	if user.Photo != "" {
		splittedStr := strings.Split(user.Photo, "/")
		oldObj = splittedStr[len(splittedStr)-1]
	}

	params := helpers.UploaderParams{
		File:         file,
		OldObjectURL: oldObj,
	}

	url, err := helpers.UploadFile(params)
	if err != nil {
		return entities.User{}, err
	}

	user.Photo = url

	result, err := s.repository.UpdateUser(user, []string{"photo"})
	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}
