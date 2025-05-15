package warehouse

type CreateWarehouseRequest struct {
	ShopId  string `json:"shopId" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
