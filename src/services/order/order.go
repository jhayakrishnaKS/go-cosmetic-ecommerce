package order

import (
	"ecommerce/src/constants"
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/services/user"
	"ecommerce/src/utils/context"
	"log"

	"github.com/google/uuid"
)

type Order struct {
	CartDAO     daos.CartDAO
	OrderDAO    daos.OrderDAO
	userService user.Users
}

func NewOrder() *Order {
	return &Order{
		CartDAO:     daos.NewCartDAO(),
		OrderDAO:    daos.NewOrderDAO(),
		userService: *user.NewUsers(),
	}
}

func (o *Order) OrderReq(req *dtos.OrderReq, ctx *context.Context, token string) *models.Order {
	userId, err := o.userService.GetAccountWithAccessToken(ctx, token)
	if err != nil {
		return nil
	}
	return &models.Order{
		ID:            uuid.New().String(),
		UserID:        userId.ID,
		ProductID:     req.ProductID,
		OrderstatusID: "dfe22fd7-ca79-40af-8475-fd06c4ed1e7e",
		Address_id:    req.AddressId,
	}
}

func (o *Order) Checkout(ctx *context.Context, req *dtos.OrderReq, token string) error {
	orderItem := o.OrderReq(req, ctx, token)

	err := o.OrderDAO.AddOrder(ctx, orderItem)
	if err != nil {
		log.Println("Unable to add item to order. Error:", err)
		return err
	}

	return o.CartDAO.RemoveAll(ctx, ctx.User.ID)
}

func (o *Order) GetAllOrder(ctx *context.Context, userId string) ([]*models.Order, error) {
	orders, err := o.OrderDAO.GetAllOrders(ctx, userId)
	if err != nil {
		log.Println("Unable to get all items in order. Error:", err)
		return nil, err
	}
	return orders, err
}

func (o *Order) UpdateStatus(ctx *context.Context, id string, updateReq *dtos.OrderReq) error {
	order, err := o.OrderDAO.GetOrderById(ctx, id)
	if err != nil {
		log.Println("Unable to find order. Error:", err)
		return constants.ErrOrderNotFound
	}
	log.Println(order)
	order.OrderstatusID = updateReq.OrderstatusID
	log.Println("order", order.OrderstatusID)
	log.Println("order", updateReq.OrderstatusID)

	err = o.OrderDAO.UpdateOrder(ctx, order)
	if err != nil {
		log.Println("Unable to update order status. Error:", err)
		return constants.ErrUnableToUpdateStatus
	}

	return nil
}
