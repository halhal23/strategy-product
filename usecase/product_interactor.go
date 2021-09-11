package usecase

import (
	"context"

	"github.com/halhal23/strategy-product/domain/model"
	"github.com/halhal23/strategy-product/domain/repository"
)

type ProductInteractor struct {
	Repo repository.IProductRepository
}

func NewProductInteractor(repo repository.IProductRepository) *ProductInteractor {
	return &ProductInteractor{
		Repo: repo,
	}
}

func (interactor *ProductInteractor) Add(ctx *context.Context, product *model.ProductModel) (int, error) {
	id, err := interactor.Repo.Save(ctx, product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (interactor *ProductInteractor) ProductByName(ctx *context.Context, name string) (*model.ProductModel, error) {
	product, err := interactor.Repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return product, nil
}