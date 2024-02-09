package controllers

import (
	"assignment-golang8-7feb/database"
	"assignment-golang8-7feb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	var order models.Order
	c.Bind(&order)

	// validasi
	needsToBeFilled := ""
	if order.CustomerName == "" {
		if needsToBeFilled != "" {
			needsToBeFilled += ", "
		}
		needsToBeFilled += "CustomerName"
	}
	for _, item := range order.Items {
		if item.ItemCode == "" {
			if needsToBeFilled != "" {
				needsToBeFilled += ", "
			}
			needsToBeFilled += "ItemCode"
		}
		if item.Description == "" {
			if needsToBeFilled != "" {
				needsToBeFilled += ", "
			}
			needsToBeFilled += "Description"
		}
		if item.Quantity == 0 {
			if needsToBeFilled != "" {
				needsToBeFilled += ", "
			}
			needsToBeFilled += "Quantity"
		}
		if needsToBeFilled != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": needsToBeFilled + " must be filled!",
			})
			return
		}
	}

	err := db.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order":   order,
		"status":  "success",
		"message": "order successfully created",
	})
}

func GetAllOrder(c *gin.Context) {
	db := database.GetDB()
	var orders []models.Order
	err := db.Model(&models.Order{}).Preload("Items").Order("id asc").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders":  orders,
		"status":  "success",
		"message": "data successfully fetched",
	})
}

func GetOrderByID(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	// validasi
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "ID missing!",
		})
		return
	}

	order := []models.Order{}
	err := db.Where("id = ?", id).Preload("Items").Find(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order":   order,
		"status":  "success",
		"message": "data successfully fetched",
	})
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var input models.Order
	c.Bind(&input)

	// validasi
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "ID missing!",
		})
		return
	}
	checkId := db.Model(&models.Order{}).Where("id = ?", id).Find(&models.Order{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Order with ID " + id + " not found",
		})
		return
	}
	if input.CustomerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "CustomerName empty!",
		})
		return
	}

	err := db.Model(&models.Order{}).Where("id = ?", id).Updates(models.Order{CustomerName: input.CustomerName}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "data successfully updated",
	})
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	// validasi
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "ID missing!",
		})
		return
	}
	checkId := db.Model(&models.Order{}).Where("id = ?", id).Find(&models.Order{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Order with ID " + id + " not found",
		})
		return
	}

	items := models.Item{}
	err := db.Where("order_id = ?", id).Delete(&items).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	order := models.Order{}
	err = db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "order successfully deleted",
	})
}
