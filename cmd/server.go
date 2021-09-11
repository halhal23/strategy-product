package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/halhal23/strategy-product/interface/rest/handler"
)

// root:@tcp(localhost:3306)/strategy_dev
func main(){
	// conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/strategy_dev")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer conn.Close()
	// var repo repository.IProductRepository = database.NewProductRepository(conn)
	// product := model.NewProductModel(0, "Ninbass2000", 19000, 1)
	// ctx := context.Background()
	// id, err := repo.Save(&ctx, product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("hello product")
	// fmt.Println(id)
	router := handler.NewRouter()
	router.Run()
}