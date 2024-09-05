package repository

import (
	"errors"

	"github.com/revandpratama/task-hub/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	r.db.Raw("SELECT email FROM users WHERE email = ?", email).Scan(&user)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *authRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user *entity.User
	r.db.Raw("SELECT username FROM users WHERE username = ?", username).Scan(&user)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
