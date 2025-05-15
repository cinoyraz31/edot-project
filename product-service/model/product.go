package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	Name        string    `gorm:"column:name;type:varchar(20)"`
	Code        string    `gorm:"column:code;type:varchar(20)"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	Status      bool      `gorm:"column:status;type:tinyint"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime"`
}
