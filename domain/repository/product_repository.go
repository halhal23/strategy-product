package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/halhal23/strategy-product/domain/model"
)

type IProductRepository interface {
	Save(*gin.Context, *model.ProductModel) (int, error)
	FindByName(ctx *gin.Context, name string) (product *model.ProductModel, err error)
}