package cart

import (
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/services/user"
	"ecommerce/src/utils/context"
	"log"

	"github.com/google/uuid"
)

type Cart struct {
	CartDAO     daos.CartDAO
	userService user.Users
}

func NewCart() *Cart {
	return &Cart{
		CartDAO:     daos.NewCartDAO(),
		userService: *user.NewUsers(),
	}
}

func (c *Cart) CreateCartReq(req *dtos.CartReq, ctx *context.Context, token string) *models.Cart {
	userId, err := c.userService.GetAccountWithAccessToken(ctx, token)
	if err != nil {
		return nil
	}
	return &models.Cart{
		ID:        uuid.New().String(),
		UserID:    userId.ID,
		ProductID: req.ProductID,
		Count:     req.Count,
	}
}

func(c *Cart)GetAllItems(ctx *context.Context, userID string)([]*models.Cart,error){
	carts,err:=c.CartDAO.GetAllItems(ctx,userID)
	if err != nil {
		log.Println("Unable to get all items in cart. Error:", err)
		return nil, err
	}
	return carts, nil
}

func (c *Cart) AddToCart(ctx *context.Context, req *dtos.CartReq, token string) error {
	cartItem := c.CreateCartReq(req, ctx, token)
	
	exists, err := c.CartDAO.CheckItemExists(ctx, req.ProductID, cartItem.UserID)
	if err != nil {
		log.Println("Unable to check if item exists in cart. Error:", err)
		return err
	}

	if exists {
		existingCartItem, err := c.CartDAO.GetCartItem(ctx, req.ProductID, cartItem.UserID)
		if err != nil {
			log.Println("Unable to get existing cart item. Error:", err)
			return err
		}
		existingCartItem.Count += req.Count
		err = c.CartDAO.UpdateItem(ctx, existingCartItem)
		if err != nil {
			log.Println("Unable to update cart item. Error:", err)
			return err
		}
	} else {
		err := c.CartDAO.AddItem(ctx, cartItem)
		if err != nil {
			log.Println("Unable to add item to cart. Error:", err)
			return err
		}
	}

	return nil
}

func (c *Cart) Delete(ctx *context.Context, id string) error {	
	err:= c.CartDAO.RemoveItem(ctx, id)
	if err != nil {
		log.Println("Unable to delete product. Error:", err)
		return err
	}
	return nil
}
