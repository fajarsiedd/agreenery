package middlewares

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

func (c *JWTConfig) NewJWTConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(c.SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return base.ErrorResponse(c, err)
			}

			return nil
		},
	}
}

func (c *JWTConfig) GenerateToken(userID, role string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(c.ExpiresDuration))))

	claims := &JWTCustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(c.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *JWTConfig) GenerateRefreshToken(userID, role string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(c.ExpiresDuration))))

	claims := &JWTCustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
			ID:        uuid.New().String(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(c.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetCurrentToken(c echo.Context) (*JWTCustomClaims, string, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, "", constants.ErrInvalidToken
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, user.Raw, nil
}
