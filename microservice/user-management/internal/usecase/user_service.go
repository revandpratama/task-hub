package usecase

import (
	"strings"
	"user-management-service/api/dto"
	"user-management-service/internal/entity"
	"user-management-service/internal/repository"
	"user-management-service/pkg/errorhandler"
	"user-management-service/pkg/util"
)

type UserService interface {
	Login(request dto.LoginRequest) (*string, error)
	Register(request dto.RegisterRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}

}

func (s userService) Login(request dto.LoginRequest) (*string, error) {
	var userDB *entity.User
	var err error

	if s.userRepo.EmailExists("email") {

	}
	if strings.Contains(request.Credential, "@") {
		userDB, err = s.userRepo.GetByEmail(request.Credential)
		if err != nil {
			return nil, err
		}
	} else {
		userDB, err = s.userRepo.GetByUsername(request.Credential)
		if err != nil {
			return nil, err
		}
	}

	// fmt.Println("userDB pass : ", userDB.Password)
	if err := util.ValidatePassword(userDB.Password, request.Password); err != nil {
		return nil, &errorhandler.UnauthorizedErr{Message: err.Error()}
		// return nil, &errorhandler.UnauthorizedErr{Message: "invalid credentials: validate password"}
	}

	tokenString, err := util.GenerateToken(userDB.ID, userDB.Role)

	if err != nil {
		return nil, err
	}

	return tokenString, nil

}

func (s userService) Register(request dto.RegisterRequest) error {

	newUser := entity.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	}

	if request.Role == "" {
		newUser.Role = "user"
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return &errorhandler.InternalServerErr{Message: err.Error()}
	}

	newUser.Password = hashedPassword

	err = s.userRepo.Create(newUser)
	if err != nil {
		return &errorhandler.InternalServerErr{Message: err.Error()}
	}

	return nil
}
