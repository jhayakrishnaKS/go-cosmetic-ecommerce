package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type CartDAO interface {
	AddItem(ctx *context.Context, cartItem *models.Cart) error
	GetCartItem(ctx *context.Context, productID, userID string) (*models.Cart, error)
	GetAllItems(ctx *context.Context,userID string) ([]*models.Cart, error)
	UpdateItem(ctx *context.Context, cartItem *models.Cart) error
	RemoveItem(ctx *context.Context, id string) error
	RemoveAll(ctx *context.Context, userID string) error
	CheckItemExists(ctx *context.Context, productID string, userID string) (bool, error)
}

type CartDAOImpl struct {
}

func NewCartDAO() CartDAO {
	return &CartDAOImpl{}
}

func (c *CartDAOImpl) AddItem(ctx *context.Context, cartItem *models.Cart) error {
	err := ctx.DB.Table("cart").Create(cartItem).Error
	if err != nil {
		log.Println("Unable to add item to cart. Error:", err)
		return err
	}
	return nil
}

func (c *CartDAOImpl) GetCartItem(ctx *context.Context, productID, userID string) (*models.Cart, error) {
	var cartItem models.Cart
	err := ctx.DB.Table("cart").Where("product_id = ? AND user_id = ?", productID, userID).First(&cartItem).Error
	if err != nil {
		log.Println("Unable to get cart item. Error:", err)
		return nil, err
	}
	return &cartItem, nil
}

func (c *CartDAOImpl) GetAllItems(ctx *context.Context,userID string) ([]*models.Cart, error) {
	var cartItems []*models.Cart
	err := ctx.DB.Table("cart").Where("user_id = ?",userID).Find(&cartItems).Error
	if err != nil {
		log.Println("Unable to get all cart items. Error:", err)
		return nil, err
	}
	return cartItems, nil
}

func (c *CartDAOImpl) UpdateItem(ctx *context.Context, cartItem *models.Cart) error {
	err := ctx.DB.Table("cart").Save(cartItem).Error
	if err != nil {
		log.Println("Unable to update cart item. Error:", err)
		return err
	}
	return nil
}

func (c *CartDAOImpl) RemoveItem(ctx *context.Context, id string) error {
	err := ctx.DB.Table("cart").Where("id = ?", id).Delete(&models.Cart{}).Error
	if err != nil {
		log.Println("Unable to remove item from cart. Error:", err)
		return err
	}
	return nil
}

func (c *CartDAOImpl) CheckItemExists(ctx *context.Context, productID, userID string) (bool, error) {
	var count int64
	err := ctx.DB.Table("cart").Where("product_id = ? AND user_id = ?", productID, userID).Count(&count).Error
	if err != nil {
		log.Println("Unable to check if item exists in cart. Error:", err)
		return false, err
	}
	return count > 0, nil
}

func (c *CartDAOImpl) RemoveAll(ctx *context.Context, userID string) error {
	err := ctx.DB.Table("cart").Where("user_id = ?", ctx.User.ID).Delete(&models.Cart{}).Error
	if err != nil {
		log.Println("Unable to remove item from cart. Error:", err)
		return err
	}
	return nil
}