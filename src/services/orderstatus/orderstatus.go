package orderstatus

import (
	"ecommerce/src/constants"
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/utils/context"

	"github.com/google/uuid"
)

type OrderStatus struct {
	orderstatusDAO daos.OrderStatusDAO
}

func NewStatus() *OrderStatus {
	return &OrderStatus{
		orderstatusDAO: daos.NewOrderStatus(),
	}
}

func (o *OrderStatus) orderStatusFromReq(req *dtos.OrderStatusReq) *models.OrderStatus {
	return &models.OrderStatus{
		ID:     uuid.New().String(),
		Status: req.Status,
	}
}

func (o *OrderStatus) CreateStatus(ctx *context.Context, req *dtos.OrderStatusReq) error {
	orderStatus := o.orderStatusFromReq(req)

	if o.orderstatusDAO.CheckStatusExists(ctx, orderStatus.Status) {
		return constants.ErrStatusTaken
	}

	return o.orderstatusDAO.Create(ctx, orderStatus)
}

