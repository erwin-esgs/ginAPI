package ProductController

import (
	"encoding/json"
	"fmt"
	"ginAPI/database"
	"ginAPI/models"
	"time"

	"github.com/gin-gonic/gin"
)

var table = "product"
var message = "message"
var dbFailMessage = "db connection failed"
var invalidInputMessage = "invalid input"
var data = "data"

func Show(c *gin.Context) {
	db, err := database.ConnectORM()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: dbFailMessage,
		})
		return
	}

	arrProduct := []models.Product{}
	db.Table(table).Find(&arrProduct).Scan(&arrProduct)
	fmt.Println(arrProduct)
	c.JSON(200, gin.H{
		data: arrProduct,
	})

}

func Detail(c *gin.Context) {
	id := c.Param("id")
	db, err := database.ConnectORM()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: dbFailMessage,
		})
		return
	}
	barang := models.Product{}
	err = db.Table(table).First(&barang, id).Error
	if err != nil {
		c.JSON(200, gin.H{
			data: nil,
		})
		return
	}
	c.JSON(200, gin.H{
		data: barang,
	})

}

func Create(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: invalidInputMessage,
		})
		return
	}
	reqBody := make(map[string]interface{})
	json.Unmarshal([]byte(jsonData), &reqBody)
	if reqBody["kuantity"] != nil && reqBody["product_name"] != nil {
		db, err := database.ConnectORM()
		if err != nil {
			c.JSON(200, gin.H{
				data:    nil,
				message: dbFailMessage,
			})
			return
		}

		barang := models.Product{}
		barang.Kuantity = reqBody["kuantity"].(string)
		barang.NamaProduct = reqBody["product_name"].(string)
		barang.CreatedAt = time.Now() //.Format("2006-02-01 15:04:05")
		barang.UpdatedAt = time.Now() //.Format("2006-02-01 15:04:05")

		result := db.Table(table).Create(&barang)
		if result.RowsAffected > 0 {
			c.JSON(200, gin.H{
				data: barang.ID,
			})
		} else {
			c.JSON(200, gin.H{
				data:    nil,
				message: "failed",
			})
		}

	} else {
		c.JSON(200, gin.H{
			data:    nil,
			message: invalidInputMessage,
		})
	}

}

func Edit(c *gin.Context) {
	id := c.Param("id")
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: invalidInputMessage,
		})
		return
	}
	reqBody := make(map[string]interface{})
	json.Unmarshal([]byte(jsonData), &reqBody)
	db, err := database.ConnectORM()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: dbFailMessage,
		})
		return
	}

	barang := models.Product{}
	err = db.Table(table).First(&barang, id).Error

	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: "not found",
		})
		return
	}

	if reqBody["kuantity"] != nil {
		barang.Kuantity = reqBody["kuantity"].(string)
	}
	if reqBody["product_name"] != nil {
		barang.NamaProduct = reqBody["product_name"].(string)
	}
	barang.UpdatedAt = time.Now() //.Format("2006-02-01 15:04:05")

	result := db.Table(table).Save(&barang)
	if result.RowsAffected > 0 {
		c.JSON(200, gin.H{
			data: result.RowsAffected,
		})
	} else {
		c.JSON(200, gin.H{
			data:    nil,
			message: "failed",
		})
	}

}

func Delete(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: invalidInputMessage,
		})
		return
	}
	reqBody := []int{}
	err = json.Unmarshal([]byte(jsonData), &reqBody)
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: invalidInputMessage,
		})
		return
	}

	db, err := database.ConnectORM()
	if err != nil {
		c.JSON(200, gin.H{
			data:    nil,
			message: dbFailMessage,
		})
		return
	}

	barang := models.Product{}
	result := db.Table(table).Delete(&barang, reqBody)
	if result.RowsAffected > 0 {
		c.JSON(200, gin.H{
			data: result.RowsAffected,
		})
	} else {
		c.JSON(200, gin.H{
			data:    nil,
			message: "failed",
		})
	}
}
