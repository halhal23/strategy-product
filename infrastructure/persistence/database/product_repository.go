package database

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/halhal23/strategy-product/domain/model"
	"github.com/halhal23/strategy-product/domain/repository"
)

type ProductRepository struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) repository.IProductRepository {
	return &ProductRepository{
		Conn: conn,
	}
}

func (repo *ProductRepository) Save(ctx *gin.Context, product *model.ProductModel) (int, error) {
	result, err := repo.Conn.Exec("INSERT INTO products (name, price, user_id) VALUES (?, ?, ?)", product.Name, product.Price, product.UserId)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	return int(id), nil
}

func (repo *ProductRepository) FindAll(ctx *gin.Context) ([]*model.ProductModel, error) {
	rows, err := repo.Conn.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*model.ProductModel
	for rows.Next() {
		var product model.ProductModel
		rows.Scan(&product.ID, &product.Name, &product.Price, &product.UserId)
		products = append(products, &product)
	}
	return products, nil
}

func (repo *ProductRepository) FindByName(ctx *gin.Context, productName string) (*model.ProductModel, error) {
	rows, err := repo.Conn.Query("SELECT * FROM products WHERE name = ?", productName)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows.Next()
	product := model.ProductModel{}
	rows.Scan(&product.ID, &product.Name, &product.Price, &product.UserId)
	return &product, nil
}