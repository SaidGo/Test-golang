package userService

import (
	"time"

	"github.com/SaidGo/Test-golang/internal/tasksService" // импортируем Task
	"gorm.io/gorm"
)

var _ = gorm.ErrInvalidData

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`

	
	Tasks []tasksService.Task `gorm:"foreignKey:UserID"`

	
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
