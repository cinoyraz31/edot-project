package warehouse_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"time"
)

type ExternalWarehouseDetail struct {
	Id        uuid.UUID
	ShopId    uuid.UUID
	Name      string
	Address   string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductId uuid.UUID
	Qty       int
	LockedQty int
}

func ShowWarehouseStock(ctx *fiber.Ctx, shopId uuid.UUID, productId uuid.UUID) (ExternalWarehouseDetail, error) {
	var response ExternalWarehouseDetail
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/shop/%s/product/%s/warehouse", os.Getenv("WAREHOUSE_URL"), shopId, productId), nil)
	if err != nil {
		return response, err
	}
	req.Header.Set("Authorization", ctx.Get("Authorization"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusOK {
		return response, errors.New("product not found")
	}

	_ = json.Unmarshal(body, &response)
	return response, nil
}
