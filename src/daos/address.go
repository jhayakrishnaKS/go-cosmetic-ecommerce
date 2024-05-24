package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type AddressDAO interface {
	AddAddress(ctx *context.Context, address *models.Address) error
	GetAddress(ctx *context.Context, userID string) (*models.Address, error)
	GetAllAddress(ctx *context.Context,userID string) ([]*models.Address, error)
	UpdateAddress(ctx *context.Context, address *models.Address) error
	RemoveAddress(ctx *context.Context, id string, userID string) error
	CheckAddressExists(ctx *context.Context, userID string) (bool, error)
}

type AddressDAOImpl struct{}

func NewAddressDAO() AddressDAO {
	return &AddressDAOImpl{}
}

func (a *AddressDAOImpl) AddAddress(ctx *context.Context, address *models.Address) error {
	err := ctx.DB.Table("address").Create(address).Error
	if err != nil {
		log.Println("Unable to add address. Error:", err)
		return err
	}
	return nil
}

func (a *AddressDAOImpl) GetAddress(ctx *context.Context, userID string) (*models.Address, error) {
	address := &models.Address{}
	err := ctx.DB.Table("address").Where("user_id = ?", userID).First(address).Error
	if err != nil {
		log.Println("Unable to get address. Error:", err)
		return nil, err
	}
	return address, nil
}

func (a *AddressDAOImpl) GetAllAddress(ctx *context.Context,userID string) ([]*models.Address, error) {
	var addresses []*models.Address
	err := ctx.DB.Table("address").Where("user_id = ?",userID).Find(&addresses).Error
	if err != nil {
		log.Println("Unable to get all addresses. Error:", err)
		return nil, err
	}
	return addresses, nil
}

func (a *AddressDAOImpl) UpdateAddress(ctx *context.Context, address *models.Address) error {
	err := ctx.DB.Table("address").Save(address).Error
	if err != nil {
		log.Println("Unable to update address. Error:", err)
		return err
	}
	return nil
}

func (a *AddressDAOImpl) RemoveAddress(ctx *context.Context, id string, userID string) error {
	err := ctx.DB.Table("address").Where("id = ? AND user_id = ?", id, userID).Delete(&models.Address{}).Error
	if err != nil {
		log.Println("Unable to remove address. Error:", err)
		return err
	}
	return nil
}

func (a *AddressDAOImpl) CheckAddressExists(ctx *context.Context, userID string) (bool, error) {
	var count int64
	err := ctx.DB.Table("address").Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		log.Println("Error checking address existence. Error:", err)
		return false, err
	}
	return count > 0, nil
}
