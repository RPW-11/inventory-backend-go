package domain

import "github.com/golang-jwt/jwt/v5"

type JwtCustomAccessClaims struct {
	FullName string `json:"full_name"`
	ID       string `json:"id"`
	Position string `json:"position"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
