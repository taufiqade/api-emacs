package service

import (
	"api-survey-go/internal/entity"
	"api-survey-go/internal/model"
	"api-survey-go/internal/model/converter"
	"api-survey-go/internal/repository"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Log      *logrus.Logger
	UserRepo *repository.UserRepository
	Validate *validator.Validate
	DB       *gorm.DB
}

func NewUserService(log *logrus.Logger, userRepo *repository.UserRepository, validate *validator.Validate, db *gorm.DB) *UserService {
	return &UserService{
		Log:      log,
		UserRepo: userRepo,
		Validate: validate,
		DB:       db,
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) Login(request *model.UserLoginRequest) (*model.LoginResponse, error) {
	user, err := s.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(request.Password, user.Password) {
		return nil, fmt.Errorf("invalid credential")
	}

	exp := time.Now().Add(time.Hour * 72).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["user_id"] = user.ID
	claims["role_id"] = user.RoleID
	claims["client_id"] = user.ClientID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     t,
		ExpiredAt: time.Unix(exp, 0),
	}, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	var user = new(entity.User)
	err := s.UserRepo.GetUserByEmail(user, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserById(id string) (*entity.User, error) {
	var user = new(entity.User)
	err := s.UserRepo.GetUserById(user, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Index() string {
	return "hello, world!"
}

func (s *UserService) Create(ctx context.Context, request *model.CreateUserRequest) (*model.User, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	err := s.UserRepo.GetUserByEmail(user, request.Email)
	if err == nil {
		return nil, fmt.Errorf("user %s already exists", user.Email)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		Email:       request.Email,
		Password:    string(password),
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		RoleID:      request.RoleID,
		ClientID:    request.ClientID,
	}

	if err = s.UserRepo.Create(tx, newUser); err != nil {
		s.Log.Warnf("Failed to create new user: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.UserToModel(newUser), nil

}

func (s *UserService) Update(ctx context.Context, request *model.UpdateUserRequest, user *entity.User) (*model.User, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		s.Log.Infof("password updated for user_id:%+v", user.ID)
		user.Password = string(password)
	}

	if request.Name != "" {
		user.Name = string(request.Name)
	}

	if request.Email != "" {
		user.Email = string(request.Email)
	}

	if request.PhoneNumber != "" {
		user.PhoneNumber = string(request.PhoneNumber)
	}

	if err := s.UserRepo.Update(tx, user); err != nil {
		s.Log.Warnf("Failed to update user: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.UserToModel(user), nil
}
