package main

import (
	"user-service/config"
	"user-service/routes"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)


func main(){
	config.ConnectDB()

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run(":8080")

}


