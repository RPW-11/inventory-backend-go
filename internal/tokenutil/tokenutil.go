package tokenutil

import (
	"fmt"
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

func IsAuthorized(requestToken, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractPositionIDFromToken(requestToken, secret string) (string, string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", "", fmt.Errorf("invalid Token")
	}

	// case for refresh token that doesn't not have "position" attribute in its claim
	if _, ok := claims["position"]; !ok {
		return claims["id"].(string), "", nil
	}

	return claims["id"].(string), claims["position"].(string), nil
}
