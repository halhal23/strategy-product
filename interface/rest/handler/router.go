package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/strategy_dev")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	handler := NewProductHandler(conn)
	router.GET("/", func(c *gin.Context) { c.JSON(200, "ok desune")})
	router.GET("/products/:name",func(c *gin.Context) { handler.Show(c)})
	router.POST("/products",func(c *gin.Context) { handler.Create(c)})
	return router
}