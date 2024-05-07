package main

import (
	"context"
	"fmt"

	"go-fundamental/projectGO/Projectuser/internal/user"
	"go-fundamental/projectGO/Projectuser/pkg/bootstrap"

	"log"
	"net/http"
)





func main() {

	server := http.NewServeMux()

	db := bootstrap.NewDB()
	logger := bootstrap.NewLogger()

	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()



	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
	
}


