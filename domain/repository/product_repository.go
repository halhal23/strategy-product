package repository

import (
	"context"

	"github.com/halhal23/strategy-product/domain/model"
)

type IProductRepository interface {
	Save(ctx *context.Context, product *model.ProductModel) (id int, err error)
	FindByName(ctx *context.Context, name string) (product *model.ProductModel, err error)
}