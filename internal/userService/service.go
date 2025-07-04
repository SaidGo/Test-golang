package userService

import (
	"github.com/SaidGo/Test-golang/internal/tasksService"
)

type Service interface {
	GetUsers() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
	GetTasksForUser(userID uint) ([]tasksService.Task, error)
}

type userServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *userServiceImpl) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *userServiceImpl) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userServiceImpl) GetTasksForUser(userID uint) ([]tasksService.Task, error) {
	user, err := s.repo.GetWithTasks(userID)
	if err != nil {
		return nil, err
	}
	return user.Tasks, nil
}
