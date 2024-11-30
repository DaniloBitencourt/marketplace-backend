package routes

import (
	"marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Registrar as rotas do CRUD de Usuario
func RegisterUsuarioRoutes(r *gin.Engine, db *gorm.DB) {
	// Criar um usuário
	r.POST("/usuarios", func(c *gin.Context) {
		var usuario models.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, usuario)
	})

	// Listar todos os usuários
	r.GET("/usuarios", func(c *gin.Context) {
		var usuarios []models.Usuario
		if err := db.Find(&usuarios).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, usuarios)
	})

	// Buscar um usuário por ID
	r.GET("/usuarios/:id", func(c *gin.Context) {
		var usuario models.Usuario
		id := c.Param("id")
		if err := db.First(&usuario, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusOK, usuario)
	})

	// Atualizar um usuário
	r.PUT("/usuarios/:id", func(c *gin.Context) {
		var usuario models.Usuario
		id := c.Param("id")
		if err := db.First(&usuario, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, usuario)
	})

	// Deletar um usuário
	r.DELETE("/usuarios/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Usuario{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})
}
