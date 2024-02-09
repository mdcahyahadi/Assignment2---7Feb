package controllers

import (
	"assignment-golang8-7feb/database"
	"assignment-golang8-7feb/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	db := database.GetDB()
	var item models.Item
	c.Bind(&item)

	// validasi
	needsToBeFilled := ""
	if item.OrderID == 0 {
		if needsToBeFilled != "" {
			needsToBeFilled += ", "
		}
		needsToBeFilled += "OrderID"
	}
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
	checkId := db.Model(&models.Order{}).Where("id = ?", item.OrderID).Find(&models.Order{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Order with ID " + strconv.FormatUint(uint64(item.OrderID), 10) + " not found",
		})
		return
	}

	err := db.Create(&item).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"item":    item,
		"status":  "success",
		"message": "item successfully added to the order",
	})
}

func GetAllItemByOrderID(c *gin.Context) {
	db := database.GetDB()
	orderID := c.Param("id")

	// validasi
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Order ID missing!",
		})
		return
	}
	checkId := db.Model(&models.Order{}).Where("id = ?", orderID).Find(&models.Order{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Order with ID " + orderID + " not found",
		})
		return
	}

	items := []models.Item{}
	err := db.Where("order_id = ?", orderID).Find(&items).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders":  items,
		"status":  "success",
		"message": "data successfully fetched",
	})
}

func GetItemByID(c *gin.Context) {
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

	items := []models.Item{}
	err := db.Where("id = ?", id).Find(&items).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders":  items,
		"status":  "success",
		"message": "data successfully fetched",
	})
}

func UpdateItem(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var input models.Item
	c.Bind(&input)

	// validasi
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "ID missing!",
		})
		return
	}
	checkId := db.Model(&models.Item{}).Where("id = ?", id).Find(&models.Item{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Item with ID " + id + " not found",
		})
		return
	}
	// needsToBeFilled := ""
	// if input.ItemCode == "" {
	// 	if needsToBeFilled != "" {
	// 		needsToBeFilled += ", "
	// 	}
	// 	needsToBeFilled += "ItemCode"
	// }
	// if input.Description == "" {
	// 	if needsToBeFilled != "" {
	// 		needsToBeFilled += ", "
	// 	}
	// 	needsToBeFilled += "Description"
	// }
	// if input.Quantity == 0 {
	// 	if needsToBeFilled != "" {
	// 		needsToBeFilled += ", "
	// 	}
	// 	needsToBeFilled += "Quantity"
	// }
	// if needsToBeFilled != "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "fail",
	// 		"message": needsToBeFilled + " empty!",
	// 	})
	// 	return
	// }
	inputField := 0
	if input.ItemCode != "" {
		inputField++
	}
	if input.Description != "" {
		inputField++
	}
	if input.Quantity != 0 {
		inputField++
	}
	if inputField == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Please at least fill one column to be updated!",
		})
		return
	}

	err := db.Model(&models.Item{}).Where("id = ?", id).Updates(models.Item{ItemCode: input.ItemCode, Description: input.Description, Quantity: input.Quantity}).Error
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

func DeleteItem(c *gin.Context) {
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
	checkId := db.Model(&models.Item{}).Where("id = ?", id).Find(&models.Item{})
	if checkId.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Item with ID " + id + " not found",
		})
		return
	}

	item := models.Item{}
	err := db.Where("id = ?", id).Delete(&item).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "item successfully deleted from the order",
	})
}
