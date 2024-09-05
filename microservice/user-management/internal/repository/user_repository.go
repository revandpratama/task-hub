package repository

import (
	"user-management-service/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	EmailExists(email string) bool
	Create(newUser entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) EmailExists(email string) bool {
	var count int64
	r.db.Raw("SELECT count(email) from users WHERE email = ?", email).Count(&count)

	return count > 0

}

func (r userRepository) Create(newUser entity.User) error {

	err := r.db.Create(&newUser).Error

	return err

}

func (r userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user).Error
	// if user.ID == 0 {
	// 	return nil, &errorhandler.NotFoundErr{Message:"user not found"}
	// }
	return &user, err
}

func (r userRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error
	// if user.ID == 0 {
	// 	return nil, &errorhandler.NotFoundErr{Message:"user not found"}
	// }
	return &user, err
}


