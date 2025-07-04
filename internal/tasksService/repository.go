package tasksService

import (
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) Repository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Create(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetByID(id uint) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Update(task Task) (Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Delete(id uint) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
