package userService

type Service interface {
	GetUsers() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &userService{repo: repo}
}

func (s *userService) GetUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *userService) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
