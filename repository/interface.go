package repository

import (
	"LoginUser/models"
	"context"
)

type Repository interface {
	UserAlreadyRegistered(ctx context.Context, username string) (bool, error)
	RegisterNewUser(ctx context.Context, userDetails models.RegisterRequest) error
	GetUserDetails(ctx context.Context, userName string) (models.UserData, error)
}
