package stock

import "github.com/google/uuid"

type OrderStockRequest struct {
	WarehouseId uuid.UUID `json:"warehouseId" validate:"required"`
	ProductId   uuid.UUID `json:"productId" validate:"required"`
	Qty         int       `json:"qty" validate:"required,numeric"`
}
