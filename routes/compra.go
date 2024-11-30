package routes

import (
	"marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCompraRoutes(r *gin.Engine, db *gorm.DB) {
	// Criar uma compra
	r.POST("/compras", func(c *gin.Context) {
		var compra models.Compra
		if err := c.ShouldBindJSON(&compra); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&compra).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, compra)
	})

	// Listar todas as compras
	r.GET("/compras", func(c *gin.Context) {
		var compras []models.Compra
		if err := db.Find(&compras).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, compras)
	})

	// Buscar compra por ID
	r.GET("/compras/:id", func(c *gin.Context) {
		var compra models.Compra
		id := c.Param("id")
		if err := db.First(&compra, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Compra não encontrada"})
			return
		}
		c.JSON(http.StatusOK, compra)
	})

	// Atualizar uma compra
	r.PUT("/compras/:id", func(c *gin.Context) {
		var compra models.Compra
		id := c.Param("id")
		if err := db.First(&compra, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Compra não encontrada"})
			return
		}
		if err := c.ShouldBindJSON(&compra); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&compra).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, compra)
	})

	// Deletar uma compra
	r.DELETE("/compras/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Compra{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Compra não encontrada"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Compra deletada com sucesso"})
	})
}
