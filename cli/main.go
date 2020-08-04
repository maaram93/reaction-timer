package main

import (
	"flag"
	"fmt"

	grpcsetup "github.com/maaram93/client/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":9092", "address where you can connect with highscore service")
	flag.Parse()

	server := grpcsetup.NewServer(*addressPtr)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start grpc server")
	}
}
