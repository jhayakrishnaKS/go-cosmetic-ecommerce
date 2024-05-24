package products

import (
	"ecommerce/src/constants"
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/utils/context"
	"log"

	"github.com/google/uuid"
)

type Product struct {
	ProductDAO daos.ProductDAO
}

func NewProduct() *Product {
	return &Product{
		ProductDAO: daos.NewProductDAO(),
	}
}

func (p *Product) CreateProductReq(req *dtos.ProductReq) *models.Products {
	return &models.Products{
		ID:            uuid.New().String(),
		Product_title: req.Product_title,
		Description:   req.Description,
		Price:         req.Price,
		Brand:         req.Brand,
	}
}


func (p *Product) CreateProduct(ctx *context.Context, req *dtos.ProductReq) error {
	if req.Product_title == "" || req.Description == "" || req.Brand == "" {
		return constants.ErrProductDescriptionEmpty
	}
	if ok, _ := p.ProductDAO.CheckTitleExists(ctx, req.Product_title); ok {
		return constants.ErrTitleTaken
	}
	product := p.CreateProductReq(req)

	return p.ProductDAO.Create(ctx, product)
}

func (p *Product) GetAllProducts(ctx *context.Context) ([]models.Products, error) {
	productPointers, err := p.ProductDAO.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var products []models.Products
	for _, ptr := range productPointers {
		products = append(products, *ptr)
	}
	return products, nil
}

func (p *Product) Delete(ctx *context.Context, id string) error {
	err := p.ProductDAO.Delete(ctx, id)
	if err != nil {
		log.Println("Unable to delete product. Error:", err)
		return err
	}
	return nil
}

func (p *Product) Update(ctx *context.Context, id string, updateReq *dtos.UpdateProductReq) error {
	product, err := p.ProductDAO.Get(ctx, id)
	if err != nil {
		log.Println("Unable to find product. Error:", err)
		return err
	}

	if updateReq.Product_title == "" || updateReq.Description == "" || updateReq.Brand == "" {
		return constants.ErrProductDescriptionEmpty
	}

	product.Product_title = updateReq.Product_title
	product.Description = updateReq.Description
	product.Brand = updateReq.Brand

	err = p.ProductDAO.Update(ctx, product)
	if err != nil {
		log.Println("Unable to update product. Error:", err)
		return err
	}

	return nil
}
