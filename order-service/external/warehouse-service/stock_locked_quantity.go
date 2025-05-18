package warehouse_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type StockLockedQuantityRequest struct {
	WarehouseId uuid.UUID `json:"warehouseId"`
	ProductId   uuid.UUID `json:"productId"`
	Qty         int       `json:"qty"`
}

func StockLockedQuantity(ctx *fiber.Ctx, warehouseId uuid.UUID, productId uuid.UUID, qty int) error {
	bodyPayload := StockLockedQuantityRequest{
		WarehouseId: warehouseId,
		ProductId:   productId,
		Qty:         qty,
	}

	requestBodyBytes, err := json.Marshal(bodyPayload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/warehouses/user/warehouse-stock/order", os.Getenv("API_GETAWAY_URL")), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ctx.Get("Authorization"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("stock not found")
	}

	return nil
}
