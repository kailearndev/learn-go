package user

import (
	"errors"
	"kai-shop-be/pkg/jwt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(req RegisterUserDTO) (User, error)
	GetUserByID(id uuid.UUID) (User, error)
	UpdateUser(id uuid.UUID, req RegisterUserDTO) (User, error)
	DeleteUser(id uuid.UUID) error
	LoginUser(req LoginUserDTO) (string, error)
	ListUsers(limit, offset int) ([]User, int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetUserByID(id uuid.UUID) (User, error) {
	return s.repo.FindByID(id)
}

func (s *service) ListUsers(limit, offset int) ([]User, int64, error) {
	items, err := s.repo.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (s *service) RegisterUser(req RegisterUserDTO) (User, error) {
	existingUser, _ := s.repo.FindByEmail(req.Email)

	if existingUser.ID != uuid.Nil {
		return User{}, errors.New("email already exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hash),
		FullName:  req.FullName,
		Role:      req.Role,
		AvatarURL: req.AvatarURL,
	}
	if err := s.repo.CreateUser(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) LoginUser(req LoginUserDTO) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}
	// Here you would normally generate a JWT or session token
	token, err := jwt.GenerateToken(user.ID.String(), user.Email, time.Hour*24)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *service) UpdateUser(id uuid.UUID, req RegisterUserDTO) (User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return User{}, err
	}
	user.Username = req.Username
	user.Email = req.Email
	user.FullName = req.FullName
	user.Role = req.Role
	user.AvatarURL = req.AvatarURL
	if req.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
	}
	if err := s.repo.UpdateUser(&user); err != nil {
		return User{}, err
	}
	return user, nil
}
func (s *service) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
func (s *service) CountProducts() (int64, error) {
	return s.repo.Count()
}
