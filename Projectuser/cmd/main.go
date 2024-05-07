package main

import (
	"fmt"
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"go-fundamental/projectGO/Projectuser/internal/user"
	"log"
	"net/http"
)





func main() {

	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "paco@gmail.com"
			},
			{
				ID:        2,
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "paco@gmail.com"
			},
		}
		}


	server.HandleFunc("/users", UserServer)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
	
}


