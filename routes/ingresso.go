package routes

import (
	"marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Registrar as rotas do CRUD de Ingresso
func RegisterIngressoRoutes(r *gin.Engine, db *gorm.DB) {
	// Criar um ingresso
	r.POST("/ingressos", func(c *gin.Context) {
		var ingresso models.Ingresso
		if err := c.ShouldBindJSON(&ingresso); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&ingresso).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, ingresso)
	})

	// Listar todos os ingressos
	r.GET("/ingressos", func(c *gin.Context) {
		var ingressos []models.Ingresso
		if err := db.Find(&ingressos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ingressos)
	})

	// Buscar ingresso por ID
	r.GET("/ingressos/:id", func(c *gin.Context) {
		var ingresso models.Ingresso
		id := c.Param("id")
		if err := db.First(&ingresso, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ingresso não encontrado"})
			return
		}
		c.JSON(http.StatusOK, ingresso)
	})

	// Atualizar um ingresso
	r.PUT("/ingressos/:id", func(c *gin.Context) {
		var ingresso models.Ingresso
		id := c.Param("id")
		if err := db.First(&ingresso, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ingresso não encontrado"})
			return
		}
		if err := c.ShouldBindJSON(&ingresso); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&ingresso).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ingresso)
	})

	// Deletar um ingresso
	r.DELETE("/ingressos/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Ingresso{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ingresso não encontrado"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Ingresso deletado com sucesso"})
	})
}
