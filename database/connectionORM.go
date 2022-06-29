package database

import (
	"ginAPI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var message = "message"
var dbFailMessage = "db connection failed"

func ConnectORM() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/testing?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Print(err.Error())
		}
	*/
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migration(c *gin.Context) {

	// var models = []interface{}{&Username{}, &Barang{}}
	db, err := ConnectORM()
	if err != nil {
		c.JSON(200, gin.H{
			message: dbFailMessage,
		})
	}
	db.Table("username").AutoMigrate(&models.Username{})
	db.Table("product").AutoMigrate(&models.Product{})
	c.JSON(200, gin.H{
		message: "migrated",
	})
}
