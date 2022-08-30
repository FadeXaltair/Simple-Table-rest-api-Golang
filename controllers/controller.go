package controllers

import (
	"log"
	"net/http"
	"practiceapi/initializers"
	"practiceapi/models"

	"github.com/gin-gonic/gin"
)

// -------------------Post ---------------------//

func CreateData(c *gin.Context){

	var body struct{
     Name string
	 Class uint
	 Rollno uint
	}
    c.Bind(&body)
	post := models.Table{Name: body.Name, Class: body.Class, RollNumber: body.Rollno}

	results := initializers.DB.Create(&post)

	if results.Error != nil {
		log.Println("error in sending data")
	}

	c.JSON(http.StatusOK, gin.H{
		"tables": post,
	})
}

// -----------------------------get ---------------------------//

func GetData(c *gin.Context) {
	var tables[] models.Table

	initializers.DB.Find(&tables)

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}


//-------------------------get one request -----------------------\\

func GetSingaleData( c *gin.Context){
   id := c.Param("id")
 
   var table models.Table
   initializers.DB.Find(&table, id)

   c.JSON(http.StatusOK, gin.H{
	    "Table": table,
   })
}


//--------------- UPdate request ---------------------\\

func UpdateReq(c *gin.Context){
	id := c.Param("id")

	var body struct {
		Name 	string
		Class 	uint
		Rollno 	uint
	}

	c.Bind(&body)

	var table models.Table
	initializers.DB.First(&table, id)

	initializers.DB.Model(&table).Updates(models.Table{Name: body.Name, Class: body.Class, RollNumber: body.Rollno})

	c.JSON(http.StatusOK, gin.H{
		"table": table,
	})
}