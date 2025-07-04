package userService

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
	GetWithTasks(id uint) (User, error) // 👈 добавлено
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *userRepository) GetWithTasks(id uint) (User, error) {
	var user User
	err := r.db.Preload("Tasks").First(&user, id).Error
	return user, err
}
