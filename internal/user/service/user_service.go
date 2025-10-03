package service

import (
	"github.com/makonheimak/user-service/internal/user/orm"
	"github.com/makonheimak/user-service/internal/user/repository"
)

type UserService interface {
	PostUser(req orm.User) (orm.User, error)
	GetAllUsers() ([]orm.User, error)
	GetUserByID(id int64) (orm.User, error)
	PatchUserByID(id int64, email, password string) (orm.User, error)
	DeleteUserByID(id int64) error
}

type Service struct {
	repo repository.UserRepository
}

func NewService(r repository.UserRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) PostUser(req orm.User) (orm.User, error) {
	err := s.repo.PostUser(&req)
	if err != nil {
		return orm.User{}, err
	}
	return req, nil
}

func (s *Service) GetAllUsers() ([]orm.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUserByID(id int64) (orm.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return orm.User{}, err
	}
	return user, nil
}

func (s *Service) PatchUserByID(id int64, email, password string) (orm.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return orm.User{}, err
	}

	user.Email = email
	user.Password = password

	if err := s.repo.PatchUserByID(&user); err != nil {
		return orm.User{}, err
	}

	return user, nil
}

func (s *Service) DeleteUserByID(id int64) error {
	err := s.repo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
