package routes

import (
	"github.com/gin-gonic/gin"
	"test/book/handlers"
	"test/book/service"
)

func RegisterBookRoutes(r *gin.Engine) {

	userHandler := handlers.NewUserHandler(service.NewUserService())
	userBookGroup := r.Group("/api/user")
	{
		userBookGroup.GET("/:id/books", userHandler.GetUserBooks)
	}
}
