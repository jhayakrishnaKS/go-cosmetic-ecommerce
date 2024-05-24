package daos

import (
	"ecommerce/src/database/models"
	"ecommerce/src/utils/context"
	"log"
)

type RefreshTokenDAO interface {
	Create(ctx *context.Context, refreshToken *models.RefreshToken) error
	Upsert(ctx *context.Context, refreshToken *models.RefreshToken) error
	Get(ctx *context.Context, token string) (*models.RefreshToken, error)
	Delete(ctx *context.Context, token string) error
}

func NewRefreshToken() RefreshTokenDAO {
	return &RefreshToken{}
}

type RefreshToken struct {
}

func (a *RefreshToken) Create(ctx *context.Context, refreshToken *models.RefreshToken) error {
	err := ctx.DB.Table("refresh_tokens").Create(refreshToken).Error
	if err != nil {
		log.Println("Unable to create RefreshToken. Err:", err)
		return err
	}

	return nil
}

func (a *RefreshToken) Upsert(ctx *context.Context, refreshToken *models.RefreshToken) error {
	err := ctx.DB.Table("refresh_tokens").Save(refreshToken).Error
	if err != nil {
		log.Println("Unable to create RefreshToken. Err:", err)
		return err
	}

	return nil
}

func (a *RefreshToken) Get(ctx *context.Context, token string) (*models.RefreshToken, error) {
	refreshtoken := &models.RefreshToken{}
	err := ctx.DB.Table("refresh_tokens").First(refreshtoken, "token=?", token).Error
	if err != nil {
		log.Println("Unable to get RefreshToken. Err:", err)
		return nil, err
	}

	return refreshtoken, nil
}

func (a *RefreshToken) Delete(ctx *context.Context, token string) error {
	err := ctx.DB.Table("refresh_tokens").Delete(&models.RefreshToken{
		Token: token,
	}).Error
	if err != nil {
		log.Println("Unable to delete Token. Err:", err)
		return err
	}

	return nil
}
