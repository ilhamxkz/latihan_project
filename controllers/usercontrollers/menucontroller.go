package usercontroller

import (
	"net/http"
	"go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexMenu(c *gin.Context) {
	var menus []models.Menu
	models.DB.Find(&menus)
	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func ShowMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if err := models.DB.First(&menu, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Menu tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&menu)
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&menu).Where("id = ?", id).Updates(&menu).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa update menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu berhasil diperbarui"})
}

func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if models.DB.Delete(&menu, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu berhasil dihapus"})
}
