package routes

import (
	"marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Registrar as rotas do CRUD de Evento
func RegisterEventoRoutes(r *gin.Engine, db *gorm.DB) {
	// Criar um evento
	r.POST("/eventos", func(c *gin.Context) {
		var evento models.Evento
		if err := c.ShouldBindJSON(&evento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&evento).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, evento)
	})

	// Listar todos os eventos
	r.GET("/eventos", func(c *gin.Context) {
		var eventos []models.Evento
		if err := db.Find(&eventos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, eventos)
	})

	// Buscar evento por ID
	r.GET("/eventos/:id", func(c *gin.Context) {
		var evento models.Evento
		id := c.Param("id")
		if err := db.First(&evento, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		c.JSON(http.StatusOK, evento)
	})

	// Atualizar um evento
	r.PUT("/eventos/:id", func(c *gin.Context) {
		var evento models.Evento
		id := c.Param("id")
		if err := db.First(&evento, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		if err := c.ShouldBindJSON(&evento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&evento).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, evento)
	})

	// Deletar um evento
	r.DELETE("/eventos/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Evento{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Evento deletado com sucesso"})
	})
}
