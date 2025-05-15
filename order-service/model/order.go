package model

import (
	"github.com/google/uuid"
	"time"
)

const STATUS_WAITING_PAYMENT = "waiting-payment"
const STATUS_PAID = "paid"
const STATUS_CANCELLED = "cancelled"

var STATUS_ALL = []string{
	STATUS_WAITING_PAYMENT,
	STATUS_PAID,
	STATUS_CANCELLED,
}

type Order struct {
	Id          uuid.UUID   `gorm:"primary_key;column:id;type:varchar(36)"`
	OrderNumber string      `gorm:"column:order_number;unique;type:varchar(36)"`
	UserId      uuid.UUID   `gorm:"column:user_id;type:varchar(36)"`
	TotalAmount float64     `gorm:"column:total_amount;type:decimal(12,2)"`
	Status      string      `gorm:"column:status;type:varchar(16)"`
	CreatedAt   time.Time   `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;type:datetime"`
	ShopOrders  []ShopOrder `gorm:"foreignkey:OrderId"`
}
