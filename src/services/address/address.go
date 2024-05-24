package address

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

type Address struct {
	AddressDAO  daos.AddressDAO
	userService user.Users
}

func NewAddress() *Address {
	return &Address{
		AddressDAO:  daos.NewAddressDAO(),
		userService: *user.NewUsers(),
	}
}

func (a *Address) AddAddressReq(req *dtos.AddressReq, ctx *context.Context, token string) *models.Address {
	userId, err := a.userService.GetAccountWithAccessToken(ctx, token)
	if err != nil {
		return nil
	}
	return &models.Address{
		ID:      uuid.New().String(),
		UserID:  userId.ID,
		DoorNo:  req.DoorNo,
		Street:  req.Street,
		City:    req.City,
		Zipcode: req.Zipcode,
	}
}

func (a *Address) GetAllAddress(ctx *context.Context, userID string) ([]*models.Address, error) {
	addresses, err := a.AddressDAO.GetAllAddress(ctx, userID)
	if err != nil {
		log.Println("Unable to get all addresses. Error:", err)
		return nil, err
	}
	return addresses, nil
}

func (a *Address) AddAddress(ctx *context.Context, req *dtos.AddressReq, token string) error {
	address := a.AddAddressReq(req, ctx, token)
	exists, err := a.AddressDAO.CheckAddressExists(ctx, address.UserID)
	if err != nil {
		log.Println("Unable to check if address exists. Error:", err)
		return err
	}

	if exists {
		log.Println("The address already exists")
		return err
	} else {
		err := a.AddressDAO.AddAddress(ctx, address)
		if err != nil {
			log.Println("Unable to add address. Error:", err)
			return err
		}
	}
	return nil
}

func (a *Address) UpdateAddress(ctx *context.Context, req *dtos.UpdateAddressReq, token string) error {
	userID := ctx.User.ID
	address, err := a.AddressDAO.GetAddress(ctx, userID)
	if err != nil {
		log.Println("Unable to find address. Error:", err)
		return err
	}

	if req.DoorNo == 0 || req.City == "" || req.Street == "" || req.Zipcode == 0 {
		return constants.ErrAddressEmpty
	}

	address.DoorNo = req.DoorNo
	address.Street = req.Street
	address.City = req.City
	address.Zipcode = req.Zipcode

	err = a.AddressDAO.UpdateAddress(ctx, address)
	if err != nil {
		log.Println("Unable to update address. Error:", err)
		return err
	}

	return nil
}

func (a *Address) DeleteAddress(ctx *context.Context, id string) error {
	userID := ctx.User.ID
	exists, err := a.AddressDAO.CheckAddressExists(ctx, userID)
	if err != nil {
		log.Println("Error checking address existence. Error:", err)
		return err
	}

	if !exists {
		return constants.ErrAddressNotFound
	}

	err = a.AddressDAO.RemoveAddress(ctx,id,userID)
	if err != nil {
		log.Println("Unable to remove address. Error:", err)
		return err
	}

	return nil
}
