package repository

import (
	"LoginUser/config"
	"LoginUser/models"
	"context"
	"database/sql"
	"log"
	"math/rand"
)

type DbClient struct {
	db        *sql.DB
	appConfig config.AppConfig
}

func (dc DbClient) UserAlreadyRegistered(ctx context.Context, userName string) (bool, error) {
	var userIsRegistered models.UserData
	if err := dc.db.QueryRow("SELECT * FROM users WHERE user_name=? ", userName).Scan(&userIsRegistered); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (dc DbClient) RegisterNewUser(ctx context.Context, userDetails models.RegisterRequest) error {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=+-)*&^%$#@!")

	b := make([]rune, dc.appConfig.SecretKeyLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	query := "INSET INTO `users` (user_name, first_name, last_name, email_id, password, secret, is_verified) VALUES(?,?,?,?,?,?)"
	_, err := dc.db.ExecContext(ctx, query, userDetails.UserName, userDetails.FirstName, userDetails.LastName, userDetails.Email, userDetails.Password, string(b), false)
	if err != nil {
		return err
	}

	return nil
}

func (dc DbClient) GetUserDetails(ctx context.Context, userName string) (models.UserData, error) {
	var userData models.UserData

	if err := dc.db.QueryRow("SELECT * FROM users WHERE user_name=? ", userName).Scan(&userData); err != nil {
		if err == sql.ErrNoRows {
			return userData, nil
		}

		return userData, err
	}

	return userData, nil
}

func (dc DbClient) VerifyUser(ctx context.Context, isVerified bool, userName string) (int64, error) {
	result, err := dc.db.Exec("UPDATE users SET is_verified = ? WHERE user_name = ?", isVerified, userName)
	if err != nil {
		log.Fatal("Error while updating users database")
		return 0, err
	}

	return result.RowsAffected()
}

func NewDbClient(db *sql.DB, appConfig config.AppConfig) Repository {
	return DbClient{
		db:        db,
		appConfig: appConfig,
	}
}
