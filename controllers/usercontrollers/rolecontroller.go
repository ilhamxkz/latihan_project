package usercontroller

import (
	"net/http"
	"go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexRole(c *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func ShowRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if err := models.DB.First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Role tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&role)
	c.JSON(http.StatusOK, gin.H{"role": role})
}

func UpdateRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&role).Where("id = ?", id).Updates(&role).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa update role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role berhasil diperbarui"})
}

func DeleteRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if models.DB.Delete(&role, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role berhasil dihapus"})
}
