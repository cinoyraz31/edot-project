package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"order-service/exceptions"
	warehouse_service "order-service/external/warehouse-service"
	"order-service/helper"
	"order-service/middleware"
	"order-service/model"
	"order-service/repository"
	"order-service/web/request"
	"time"
)

type OrderControllerImpl struct {
	DB                  *gorm.DB
	OrderRepository     repository.OrderRepository
	OrderItemRepository repository.OrderItemRepository
	ShopOrderRepository repository.ShopOrderRepository
}

func NewOrderController(
	orderItemRepository repository.OrderItemRepository,
	DB *gorm.DB,
	orderRepository repository.OrderRepository,
	shopOrderRepository repository.ShopOrderRepository,
) *OrderControllerImpl {
	return &OrderControllerImpl{
		OrderItemRepository: orderItemRepository,
		DB:                  DB,
		OrderRepository:     orderRepository,
		ShopOrderRepository: shopOrderRepository,
	}
}

func (o OrderControllerImpl) Add(ctx *fiber.Ctx) error {
	var data request.OrderRequest
	claims := ctx.Locals("claims").(middleware.CheckTokenResponse)

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	if len(data.Order) > 0 {
		order := model.Order{
			Id:          uuid.New(),
			OrderNumber: helper.OrderNumber(),
			UserId:      claims.Data.Id,
			Status:      model.STATUS_WAITING_PAYMENT,
		}

		mapShopOrder := make(map[uuid.UUID]struct {
			Subtotal   float64
			OrderItems []model.OrderItem
		})

		var totalAmount float64
		for _, orderItem := range data.Order {
			//	check warehouse & stock
			warehouseResponse, err := warehouse_service.ShowWarehouseStock(ctx, orderItem.ShopId, orderItem.ProductId)
			if err != nil {
				return exceptions.ErrorHandlerBadRequest(ctx, "There a product not have shop and warehouse")
			}
			if (warehouseResponse.Qty - warehouseResponse.LockedQty) < orderItem.Quantity {
				return exceptions.ErrorHandlerBadRequest(ctx, "There a product not enough quantity")
			}

			subTotal := float64(orderItem.Quantity) * orderItem.Price
			entry := mapShopOrder[orderItem.ShopId]
			entry.Subtotal += subTotal
			entry.OrderItems = append(entry.OrderItems, model.OrderItem{
				Qty:         orderItem.Quantity,
				Price:       orderItem.Price,
				ProductId:   orderItem.ProductId,
				SubTotal:    subTotal,
				Id:          uuid.New(),
				WarehouseId: warehouseResponse.Id,
			})
			mapShopOrder[orderItem.ShopId] = entry

			totalAmount += orderItem.Price * float64(orderItem.Quantity)
		}

		order.TotalAmount = totalAmount

		for shopId, shopOrder := range mapShopOrder {
			order.ShopOrders = append(order.ShopOrders, model.ShopOrder{
				Id:         uuid.New(),
				ShopId:     shopId,
				SubTotal:   shopOrder.Subtotal,
				OrderItems: shopOrder.OrderItems,
			})
		}

		// locked qty
		for _, value := range order.ShopOrders {
			for _, orderItem := range value.OrderItems {
				if err := warehouse_service.StockLockedQuantity(ctx, orderItem.WarehouseId, orderItem.ProductId, orderItem.Qty); err != nil {
					return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
				}
			}
		}

		if err := o.OrderRepository.Create(o.DB, order); err != nil {
			return exceptions.ErrorHandlerBadRequest(ctx, "Faield to create order")
		}

		go OrderCancelled(ctx.Get("Authorization"), o, order)

		return ctx.Status(fiber.StatusCreated).JSON("")
	} else {
		return exceptions.ErrorHandlerBadRequest(ctx, "Order is empty min 1 order product")
	}
}

func OrderCancelled(authorization string, o OrderControllerImpl, order model.Order) {
	time.Sleep(1 * time.Minute) // dummy 1 menit waktu tunggu payment
	result, err := o.OrderRepository.FindBy(o.DB, map[string]interface{}{
		"order_number": order.OrderNumber,
		"status":       model.STATUS_WAITING_PAYMENT,
	})
	if err == nil {
		result.Status = model.STATUS_CANCELLED
		if err := o.OrderRepository.Update(o.DB, result); err != nil {
			log.Fatal(err.Error())
		}
		for _, shopOrders := range order.ShopOrders {
			for _, orderItem := range shopOrders.OrderItems {
				warehouse_service.StockReleaseQuantity(authorization, orderItem.WarehouseId, orderItem.ProductId, orderItem.Qty)
			}
		}
	}

}

func (o OrderControllerImpl) Show(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderControllerImpl) List(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderControllerImpl) Pay(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
