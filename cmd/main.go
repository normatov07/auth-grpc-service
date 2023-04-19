package main

import (
	"auth/api"
	db "auth/pkg/db/sqlc"
	"auth/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connection error to database: ", err)
	}
	fmt.Println("Run server main.go file !!!!")
	store := db.NewStore(conn)
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("can not create server: ", err)
	}
	server.Start(config.ServerAddress)

	//////FOR GRPC Server
	// lis, err := net.Listen("tcp", config.ServerAddress)
	// if err != nil {
	// 	log.Fatal("Can not start server: ", err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterAuthServceServer(grpcServer, server)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalln("Failed to serve: ", err)
	// }
	// fmt.Println("Serve server, address: 127.0.0.1:8080")
}
