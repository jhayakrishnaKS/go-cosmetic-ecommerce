package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type OrderStatusDAO interface {
	Create(ctx *context.Context, orderstatus *models.OrderStatus) error
	CheckStatusExists(ctx *context.Context, status string) bool
}

func NewOrderStatus() OrderStatusDAO {
	return &OrderStatus{}
}

type OrderStatus struct {
}

func (o *OrderStatus) CheckStatusExists(ctx *context.Context, status string) bool {
	var cnt int
	err := ctx.DB.Table("orderstatus").Select("count(*)").Where("status=?", status).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to check status existence. Err:", err)
		return false
	}
	return cnt > 0
}

func (o *OrderStatus) Create(ctx *context.Context, orderstatus *models.OrderStatus) error {
	err := ctx.DB.Table("orderstatus").Create(orderstatus).Error
	if err != nil {
		log.Println("Unable to create status. Err:", err)
		return err
	}
	return nil
}

func (o *OrderStatus) EditStatus(ctx *context.Context, orderstatus *models.OrderStatus) error {
	err := ctx.DB.Table("orderstatus").Save(orderstatus).Error
	if err != nil {
		log.Println("Unable to update status. Err:", err)
		return err
	}
	return nil
}
