package tasksService

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Repository interface {
	GetAll() ([]Task, error)
	Create(Task) (Task, error)
	GetByID(id uint) (Task, error)
	Update(Task) (Task, error)
	Delete(id uint) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateTask(t Task) (Task, error) {
	return s.repo.Create(t)
}

func (s *Service) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateTask(t Task) (Task, error) {
	return s.repo.Update(t)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
