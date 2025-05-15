package stock

import "github.com/google/uuid"

type CreateStockRequest struct {
	WarehouseId uuid.UUID `json:"warehouseId" validate:"required"`
	ProductId   uuid.UUID `json:"productId" validate:"required"`
	Qty         int       `json:"qty" validate:"required,numeric"`
}

type UpdateStockRequest struct {
	Qty int `json:"qty" validate:"required,numeric"`
}
