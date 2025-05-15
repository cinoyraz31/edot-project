package request

import "github.com/google/uuid"

type OrderItemRequest struct {
	ShopId    uuid.UUID `json:"shopId" validate:"required,uuid"`
	ProductId uuid.UUID `json:"productId" validate:"required,uuid"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
	Price     float64   `json:"price" validate:"required,gt=0"`
}

type OrderRequest struct {
	Order []OrderItemRequest `json:"order" validate:"required,dive,required"`
}
