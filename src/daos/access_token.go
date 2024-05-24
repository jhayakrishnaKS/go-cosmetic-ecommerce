package daos

import (
	"ecommerce/src/utils/context"
	"ecommerce/src/database/models"
	"log"
)

type AccessTokenDAO interface {
	Create(ctx *context.Context, accessToken *models.AccessToken) error
	Upsert(ctx *context.Context, accessToken *models.AccessToken) error
	Get(ctx *context.Context, token string) (*models.AccessToken, error)
	Delete(ctx *context.Context, token string) error
}

func NewAccessToken() AccessTokenDAO {
	return &AccessToken{}
}

type AccessToken struct {
}

func (a *AccessToken) Create(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_token").Create(accessToken).Error
	if err != nil {
		log.Println("Unable to create AccessToken. Err:", err)
		return err
	}

	return nil
}

func (a *AccessToken) Upsert(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_token").Save(accessToken).Error
	if err != nil {
		log.Println("Unable to create AccessToken. Err:", err)
		return err
	}

	return nil
}

func (a *AccessToken) Get(ctx *context.Context, token string) (*models.AccessToken, error) {
	accesstoken := &models.AccessToken{}
	err := ctx.DB.Table("access_token").First(accesstoken, "token=?", token).Error
	if err != nil {
		log.Println("Unable to get AccessToken. Err:", err)
		return nil, err
	}

	return accesstoken, nil
}

func (a *AccessToken) Delete(ctx *context.Context, token string) error {
	err := ctx.DB.Table("access_token").Delete(&models.AccessToken{
		Token: token,
	}).Error
	if err != nil {
		log.Println("Unable to delete Token. Err:", err)
		return err
	}

	return nil
}
