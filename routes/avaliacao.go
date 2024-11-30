package routes

import (
	"marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAvaliacaoRoutes(r *gin.Engine, db *gorm.DB) {
	// Criar uma avaliação
	r.POST("/avaliacoes", func(c *gin.Context) {
		var avaliacao models.Avaliacao
		if err := c.ShouldBindJSON(&avaliacao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&avaliacao).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, avaliacao)
	})

	// Listar todas as avaliações
	r.GET("/avaliacoes", func(c *gin.Context) {
		var avaliacoes []models.Avaliacao
		if err := db.Find(&avaliacoes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, avaliacoes)
	})

	// Buscar avaliação por ID
	r.GET("/avaliacoes/:id", func(c *gin.Context) {
		var avaliacao models.Avaliacao
		id := c.Param("id")
		if err := db.First(&avaliacao, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Avaliação não encontrada"})
			return
		}
		c.JSON(http.StatusOK, avaliacao)
	})

	// Atualizar uma avaliação
	r.PUT("/avaliacoes/:id", func(c *gin.Context) {
		var avaliacao models.Avaliacao
		id := c.Param("id")
		if err := db.First(&avaliacao, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Avaliação não encontrada"})
			return
		}
		if err := c.ShouldBindJSON(&avaliacao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&avaliacao).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, avaliacao)
	})

	// Deletar uma avaliação
	r.DELETE("/avaliacoes/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Avaliacao{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Avaliação não encontrada"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Avaliação deletada com sucesso"})
	})
}
