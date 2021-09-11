package handler

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/halhal23/strategy-product/domain/model"
	"github.com/halhal23/strategy-product/infrastructure/persistence/database"
	"github.com/halhal23/strategy-product/pkg/pb"
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

func (handler *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	input := req.GetProductInput()
	product := model.ProductModel{
		Name: input.Name,
		Price: int(input.Price),
		UserId: int(input.UserId),
	}
	var c *gin.Context
	id, err := handler.Interactor.Add(c, &product)
	if err != nil {
		return nil, err 
	} 
	return &pb.CreateProductResponse{
		Id: int64(id),
	}, nil
}

func (handler *ProductHandler) ReadByNameProduct(ctx context.Context, req *pb.ReadByNameProductRequest) (*pb.ReadByNameProductResponse, error) {
	queryName := req.GetName()
	var c *gin.Context
	product, err := handler.Interactor.ProductByName(c, queryName)
	if err != nil {
		return nil, err
	}
	return &pb.ReadByNameProductResponse{
		Product: &pb.Product{
			Id: int64(product.ID),
			Name: product.Name,
			Price: int64(product.Price),
			UserId: int64(product.UserId),
		},
	}, nil
}

func (handler *ProductHandler) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	var c *gin.Context
	products, err := handler.Interactor.Products(c)
	if err != nil {
		return nil, err
	}
	var response []*pb.Product
	for _, v := range products {
		product := pb.Product{
			Id: int64(v.ID),
			Name: v.Name,
			Price: int64(v.Price),
			UserId: int64(v.UserId),
		}
		response = append(response, &product)
	}
	return &pb.ListProductResponse{
		Product: response,
	}, nil
}