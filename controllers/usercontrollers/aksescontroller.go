package usercontroller

import (
	"net/http"
	"go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexAkses(c *gin.Context) {
	var aksesList []models.Akses
	models.DB.Find(&aksesList)
	c.JSON(http.StatusOK, gin.H{"akses": aksesList})
}

func ShowAkses(c *gin.Context) {
	var akses models.Akses
	id := c.Param("id")
	if err := models.DB.First(&akses, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Akses tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"akses": akses})
}

func CreateAkses(c *gin.Context) {
	var akses models.Akses
	if err := c.ShouldBindJSON(&akses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&akses)
	c.JSON(http.StatusOK, gin.H{"akses": akses})
}

func UpdateAkses(c *gin.Context) {
	var akses models.Akses
	id := c.Param("id")
	if err := c.ShouldBindJSON(&akses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&akses).Where("id = ?", id).Updates(&akses).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa update akses"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Akses berhasil diperbarui"})
}

func DeleteAkses(c *gin.Context) {
	var akses models.Akses
	id := c.Param("id")
	if models.DB.Delete(&akses, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus akses"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Akses berhasil dihapus"})
}
