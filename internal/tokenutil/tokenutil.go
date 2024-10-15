package tokenutil

import (
	"time"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry)))
	claims := &domain.JwtCustomAccessClaims{
		FullName: user.FullName,
		ID:       user.ID,
		Position: user.Position,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry)))
	claims := &domain.JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
