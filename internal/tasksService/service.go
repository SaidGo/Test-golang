package tasksService

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}

// Repository описывает методы доступа к данным
type Repository interface {
	GetAll() ([]Task, error)
	GetByUserID(userID uint) ([]Task, error) // 👈 добавлено
	Create(Task) (Task, error)
	GetByID(id uint) (Task, error)
	Update(Task) (Task, error)
	Delete(id uint) error
}

// Service описывает бизнес-логику
type Service interface {
	GetAllTasks() ([]Task, error)
	GetTasksByUser(userID uint) ([]Task, error) // 👈 добавлено
	CreateTask(Task) (Task, error)
	GetTaskByID(id uint) (Task, error)
	UpdateTask(Task) (Task, error)
	DeleteTask(id uint) error
}

type serviceImpl struct {
	repo Repository
}

// NewService конструктор, возвращает интерфейс Service
func NewService(r Repository) Service {
	return &serviceImpl{repo: r}
}

func (s *serviceImpl) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) GetTasksByUser(userID uint) ([]Task, error) {
	return s.repo.GetByUserID(userID)
}

func (s *serviceImpl) CreateTask(t Task) (Task, error) {
	return s.repo.Create(t)
}

func (s *serviceImpl) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetByID(id)
}

func (s *serviceImpl) UpdateTask(t Task) (Task, error) {
	return s.repo.Update(t)
}

func (s *serviceImpl) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
