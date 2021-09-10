package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/halhal23/strategy-product/domain/model"
)

type ProductRepository struct {
	SqlHandler *sql.DB
}

func (repo *ProductRepository) Save(ctx context.Context, product model.ProductModel) *model.ProductModel {
	result, err := repo.SqlHandler.Exec("INSERT INTO products (name, price, user_id) VALUES (?, ?, ?)", product.Name, product.Price, product.UserId)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &model.ProductModel{
		ID: int(id),
		Name: product.Name,
		Price: product.Price,
		UserId: product.UserId,
	}
}