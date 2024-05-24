package handlers

import (
	"ecommerce/src/dtos"
	"ecommerce/src/services/products"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Products(c *context.Context) {
	req := &dtos.ProductReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if c.User.RoleID == "947019b0-0970-4dd2-9dfa-8320cd871e06" {
		err := products.NewProduct().CreateProduct(c, req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product created successfully",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to create products",
		})
	}
}

func GetAllProducts(c *context.Context) {
	products, err := products.NewProduct().GetAllProducts(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func UpdateProducts(c *context.Context) {
    if c.User.RoleID != "947019b0-0970-4dd2-9dfa-8320cd871e06" {
        c.JSON(http.StatusForbidden, gin.H{
            "error": "You do not have permission to update products",
        })
        return
    }

    productID := c.Param("id")
    if productID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Product ID is required",
        })
        return
    }

    var updateReq dtos.UpdateProductReq
    if err := c.BindJSON(&updateReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    err := products.NewProduct().Update(c, productID, &updateReq)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "msg": "Product updated successfully",
    })
}


func DeleteProducts(c *context.Context) {
    if c.User.RoleID != "947019b0-0970-4dd2-9dfa-8320cd871e06" {
        c.JSON(http.StatusForbidden, gin.H{
            "error": "You do not have permission to delete products",
        })
        return
    }

    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID parameter is required",
        })
        return
    }

    err := products.NewProduct().Delete(c, id)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "msg": "Product deleted successfully",
    })
}

