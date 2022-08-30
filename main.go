package main

import (
	"practiceapi/controllers"
	"practiceapi/initializers"
	"practiceapi/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Table{})
	r := gin.Default()

	r.POST("/tables", controllers.CreateData)
	r.PUT("/tables/:id", controllers.UpdateReq)
	r.GET("/tables", controllers.GetData)
	r.GET("/tables/:id", controllers.GetSingaleData)
	r.Run()

}