package main

import (
	"awesomeProject/controller"
	"awesomeProject/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	service.InitializeBooks()
	service.InitializeUser()

	v1 := router.Group("/api/v1")
	{
		book := v1.Group("/books")
		{
			book.GET("/", controller.GetAllBooks)
			book.POST("/", controller.AddBook)

			book.GET("/:id", controller.GetBookByID)
			book.PUT("/:id", controller.UpdateBookByID)
			book.DELETE("/:id", controller.DeleteBookByID)
		}

		users := v1.Group("/users")
		{
			users.GET("/", controller.GetAllUsers)
			users.POST("/", controller.Register)

			users.GET("/:id", controller.Profile)
			users.PUT("/:id", controller.UpdateProfile)
			users.DELETE("/:id", controller.DeleteAccount)
		}
	}

	err := router.Run("0.0.0.0:" + os.Getenv("PORT"))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
