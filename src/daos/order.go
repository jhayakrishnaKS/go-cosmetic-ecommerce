package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type OrderDAO interface {
	AddOrder(ctx *context.Context, order *models.Order) error
	GetOrder(ctx *context.Context, productID,userID string) (*models.Order, error)
	GetAllOrders(ctx *context.Context, userID string) ([]*models.Order, error)
	RemoveOrder(ctx *context.Context, orderID string, userID string) error
	CheckOrderExists(ctx *context.Context, productID string, userID string) (bool, error)
	UpdateOrder(ctx *context.Context, order *models.Order) error
	GetOrderById(ctx *context.Context, id string) (*models.Order, error)
}
type OrderDAOImpl struct {
}

func NewOrderDAO() OrderDAO {
	return &OrderDAOImpl{}
}

func (o *OrderDAOImpl) AddOrder(ctx *context.Context, order *models.Order) error {
	err := ctx.DB.Table("order").Create(order).Error
	if err != nil {
		log.Println("Unable to add order. Error:", err)
		return err
	}
	return nil
}

func (o *OrderDAOImpl) GetOrder(ctx *context.Context, productID,userID string)  (*models.Order, error) {
	var order models.Order
	err := ctx.DB.Table("order").Where("product_id = ? AND user_id = ?", productID, userID).First(&order).Error
	if err != nil {
		log.Println("Unable to get order. Error:", err)
		return nil, err
	}
	return &order, nil
}

func (o *OrderDAOImpl) GetAllOrders(ctx *context.Context, userID string) ([]*models.Order, error) {
	var orders []*models.Order
	err := ctx.DB.Table("order").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		log.Println("Unable to get all orders. Error:", err)
		return nil, err
	}
	return orders, nil
}

func (o *OrderDAOImpl) RemoveOrder(ctx *context.Context, orderID string, userID string) error {
	err := ctx.DB.Table("order").Where("id = ? AND user_id = ?", orderID, userID).Delete(&models.Order{}).Error
	if err != nil {
		log.Println("Unable to remove order. Error:", err)
		return err
	}
	return nil
}

func (o *OrderDAOImpl) CheckOrderExists(ctx *context.Context, productID string, userID string) (bool, error) {
	var count int64
	err := ctx.DB.Table("order").Model(&models.Order{}).Where("product_id = ? AND user_id = ?", productID, userID).Count(&count).Error
	if err != nil {
		log.Println("Unable to check if order exists. Error:", err)
		return false, err
	}
	return count > 0, nil
}

func (o *OrderDAOImpl) UpdateOrder(ctx *context.Context, order *models.Order) error {
	err := ctx.DB.Table("order").Save(order).Error
	if err != nil {
		log.Println("Unable to update order. Error:", err)
		return err
	}
	return nil
}

func (o *OrderDAOImpl) GetOrderById(ctx *context.Context, id string) (*models.Order, error) {
	var order models.Order
	err := ctx.DB.Table("order").Where("id = ?", id).First(&order).Error
	if err != nil {
		log.Println("Unable to find order. Error:", err)
		return nil, err
	}
	return &order, nil
}