package daos

import (
	"ecommerce/src/utils/context"
	"ecommerce/src/database/models"
	"log"
)

type ProductDAO interface {
	Create(ctx *context.Context, product *models.Products) error
	Get(ctx *context.Context, id string) (*models.Products, error)
	GetAll(ctx *context.Context) ([]*models.Products, error)
	Update(ctx *context.Context, product *models.Products) error
	Delete(ctx *context.Context, id string) error
	CheckTitleExists(ctx *context.Context, Product_title string) (bool, error)
}

type ProductsDAO struct {
}

func NewProductDAO() ProductDAO {
	return &ProductsDAO{}
}

func (p *ProductsDAO) Create(ctx *context.Context, product *models.Products) error {
	err := ctx.DB.Table("products").Create(product).Error
	if err != nil {
		log.Println("Unable to create product. Error:", err)
		return err
	}
	return nil
}

func (p *ProductsDAO) Get(ctx *context.Context, id string) (*models.Products, error) {
    product := &models.Products{}
    err := ctx.DB.Table("products").Where("id = ?", id).First(product).Error
    if err != nil {
        log.Println("Unable to read product. Error:", err)
        return nil, err
    }
    return product, nil
}

func (p *ProductsDAO) GetAll(ctx *context.Context) ([]*models.Products, error) {
	var products []*models.Products
	err := ctx.DB.Table("products").Find(&products).Error
	if err != nil {
		log.Println("Unable to get all products. Error:", err)
		return nil, err
	}
	return products, nil
}

func (p *ProductsDAO) Update(ctx *context.Context, product *models.Products) error {
	err := ctx.DB.Table("products").Save(product).Error
	if err != nil {
		log.Println("Unable to update product. Error:", err)
		return err
	}
	return nil
}

func (p *ProductsDAO) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("products").Where("id = ?", id).Delete(&models.Products{}).Error
	if err != nil {
		log.Println("Unable to delete product. Error:", err)
		return err
	}
	return nil
}

func (p *ProductsDAO) CheckTitleExists(ctx *context.Context, Product_title string) (bool, error) {
	var cnt int
	err := ctx.DB.Table("products").Select("count(*)").Where("product_title = ?", Product_title).Scan(&cnt).Error
	if err != nil {
		log.Println("The product title already exists. Error:", err)
		return false, err
	}
	return cnt > 0, nil
}
