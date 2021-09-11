package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/halhal23/strategy-product/domain/model"
	"github.com/halhal23/strategy-product/infrastructure/persistence/database"
	"github.com/halhal23/strategy-product/usecase"
)

type ProductHandler struct {
	Interactor *usecase.ProductInteractor
}

func NewProductHandler(conn *sql.DB) *ProductHandler {
	return &ProductHandler{
		Interactor: usecase.NewProductInteractor(database.NewProductRepository(conn)),
	}
}

func (handler *ProductHandler) Index(ctx *gin.Context) {
	products, err := handler.Interactor.Products(ctx)
	if err != nil {
		ctx.JSON(500, err.Error())
	}
	ctx.JSON(200, products)
}

func (handler *ProductHandler) Show(ctx *gin.Context) {
	name := ctx.Param("name")
	product, err := handler.Interactor.ProductByName(ctx, name)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, product)
}

func (handler *ProductHandler) Create(ctx *gin.Context) {
	var product = model.ProductModel{}
	ctx.Bind(&product)
	id, err := handler.Interactor.Add(ctx, &product)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(201, id)
}