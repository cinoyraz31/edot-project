package model

import (
	"github.com/google/uuid"
	"time"
)

type Shop struct {
	Id        uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	Code      string    `gorm:"column:code;type:varchar(50)"`
	Name      string    `gorm:"column:name;type:varchar(20)"`
	Status    bool      `gorm:"column:status;type:tinyint"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
