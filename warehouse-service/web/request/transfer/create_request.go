package transfer

import "github.com/google/uuid"

type CreateTransferRequest struct {
	FromWarehouseId uuid.UUID `json:"fromWarehouseId" validate:"required"`
	ToWarehouseId   uuid.UUID `json:"toWarehouseId" validate:"required"`
	ProductId       uuid.UUID `json:"productId" validate:"required"`
	Qty             int       `json:"qty" validate:"required"`
}
