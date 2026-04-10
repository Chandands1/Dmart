package routes


import (
	"user-service/handler"
	//"grom.io/gorm"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	userHandler := handler.NewUserHandler()

	user := r.Group("/api/v1/users")
	{
		user.POST("/register", userHandler.Register)
	// 	user.POST("/login", userHandler.Login)

	// 	user.GET("/:id", userHandler.GetUserByID)

	// 	// Protected routes (later add middleware)
	// 	user.PUT("/:id", userHandler.UpdateUser)
	// 	user.DELETE("/:id", userHandler.DeleteUser)
	// 
	}
     


}


