package bootstrap

import (
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"go-fundamental/projectGO/Projectuser/internal/user"
	"log"
	"os"
)


func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		
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


}
