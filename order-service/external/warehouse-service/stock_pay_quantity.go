package warehouse_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
)

type StockPayQuantityRequest struct {
	WarehouseId uuid.UUID `json:"warehouseId"`
	ProductId   uuid.UUID `json:"productId"`
	Qty         int       `json:"qty"`
}

func StockPayQuantity(authorization string, warehouseId uuid.UUID, productId uuid.UUID, qty int) {
	bodyPayload := StockPayQuantityRequest{
		WarehouseId: warehouseId,
		ProductId:   productId,
		Qty:         qty,
	}

	requestBodyBytes, err := json.Marshal(bodyPayload)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/warehouses/user/warehouse-stock/pay", os.Getenv("API_GETAWAY_URL")), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusNoContent {
		log.Fatal(errors.New("stock not found"))
	}
}
