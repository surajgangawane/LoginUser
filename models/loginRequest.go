package models

import "github.com/golang-jwt/jwt"

type UserData struct {
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	EmailId    string `json:"email_id"`
	Password   string `json:"password"`
	SecretKey  string `json:"secret_key"`
	IsVerified string `json:"is_verified"`
}

type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	Token string
	User  struct {
		UserName  string
		FirstName string
		Email     string
		LastName  string
	}
}

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom_claims"`
}
