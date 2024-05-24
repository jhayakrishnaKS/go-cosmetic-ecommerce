package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type UserDAO interface {
	Create(ctx *context.Context, user *models.Users) error
	GetAll(ctx *context.Context) ([]*models.Users, error)
	Get(ctx *context.Context, id string) (*models.Users, error)
	CheckEmailExists(ctx *context.Context, email string) (bool, error)
	CheckUsernameExists(ctx *context.Context, username string) (bool, error)
	GetAccountFromEmailOrMobile(ctx *context.Context, email, username, roleID string) (*models.Users, error)
}

type Users struct {
}

func NewUser() UserDAO {
	return &Users{}
}

func (u *Users) Create(ctx *context.Context, user *models.Users) error {
	err := ctx.DB.Table("users").Create(user).Error
	if err != nil {
		log.Println("Unable to create user. Err:", err)
		return err
	}
	return nil
}

func (u *Users) CheckEmailExists(ctx *context.Context, email string) (bool, error) {
	var count int64
	err := ctx.DB.Table("users").Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.Println("Error checking email existence:", err)
		return false, err
	}
	return count > 0, nil
}

func (u *Users) CheckUsernameExists(ctx *context.Context, username string) (bool, error) {
	var count int64
	err := ctx.DB.Table("users").Where("username = ?", username).Count(&count).Error
	if err != nil {
		log.Println("Error checking username existence:", err)
		return false, err
	}
	return count > 0, nil
}

func (u *Users) GetAll(ctx *context.Context) ([]*models.Users, error) {
	var users []*models.Users
	err := ctx.DB.Table("users").Find(&users).Error
	if err != nil {
		log.Println("Unable to get all users. Error:", err)
		return nil, err
	}
	return users, nil
}

func(u *Users)GetAccountFromEmailOrMobile(ctx *context.Context, email, username, roleID string) (*models.Users, error){
	user:=&models.Users{}
	err:=ctx.DB.Table("users").Where("(email=? OR username=?) AND role_id=?",email,username,roleID).First(user).Error
	if err != nil {
		log.Println("Unable to find user.Err:", err)
		return nil, err
	}
	return user,nil
}

func (u *Users)Get(ctx *context.Context, id string) (*models.Users, error){
	user:=&models.Users{}
	err := ctx.DB.Table("users").First(user, "id=?", id).Error
	if err != nil {
		log.Println("Unable to create account.Err:", err)
		return nil, err
	}
	return user, nil
}

