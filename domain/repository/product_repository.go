package repository

import (
	"context"

	"github.com/halhal23/strategy-product/domain/model"
)

type IProductRepository interface {
	Save(*context.Context, *model.ProductModel) (int, error)
	FindByName(ctx *context.Context, name string) (product *model.ProductModel, err error)
}