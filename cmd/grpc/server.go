package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/halhal23/strategy-product/interface/grpc/handler"
	"github.com/halhal23/strategy-product/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err.Error())
	}

	server := grpc.NewServer()

	conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/strategy_dev")
	if err != nil {
		fmt.Println(err.Error())
	}

	handler := handler.NewProductHandler(conn)
	pb.RegisterProductServiceServer(server, handler)

	reflection.Register(server)

	fmt.Println("start grpc server")
	if err := server.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
}