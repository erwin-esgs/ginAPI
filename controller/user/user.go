package UserController

import (
	"fmt"
	"ginAPI/database"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	con := database.ConnectSQL()
	result, err := con.Exec("Select * from username ")
	if err == nil {
		defer con.Close()
		fmt.Println(result)
		c.JSON(200, gin.H{
			"message": result,
		})
	} else {
		fmt.Println(err)
	}

}

func Detail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Detail " + id,
	})
}

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create",
	})
}

func Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Edit",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
