package main

import (
	"context"
	"fmt"
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"go-fundamental/projectGO/Projectuser/internal/user"
	"log"
	"net/http"
	"os"
)





func main() {

	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "paco@gmail.com",
			},
			{
				ID:        2,
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "paco@gmail.com",
			},
		
		},

		MaxID: 2,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)

	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()



	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
	
}


