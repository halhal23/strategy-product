package handler

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/halhal23/strategy-product/domain/model"
	"github.com/halhal23/strategy-product/infrastructure/persistence/database"
	"github.com/halhal23/strategy-product/pkg/pb/pkg/pb"
	"github.com/halhal23/strategy-product/usecase"
	_ "google.golang.org/grpc"
)

type ProductHandler struct {
	Interactor *usecase.ProductInteractor
}

func NewProductHandler(conn *sql.DB) *ProductHandler {
	return &ProductHandler{
		Interactor: usecase.NewProductInteractor(database.NewProductRepository(conn)),
	}
}

func (handler *ProductHandler) CreateProduct(ctx *context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	input := req.GetProductInput()
	product := model.ProductModel{
		Name: input.Name,
		Price: int(input.Price),
		UserId: int(input.UserId),
	}
	var c *gin.Context
	id, err := handler.Interactor.Add(c, &product)
	if err == nil {
		return nil, err 
	} 
	return &pb.CreateProductResponse{
		Id: int64(id),
	}, nil
}