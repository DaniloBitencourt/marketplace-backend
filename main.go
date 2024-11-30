package main

import (
	"log"
	"marketplace-backend/database"
	"marketplace-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	r := gin.Default()

	routes.RegisterUsuarioRoutes(r, db)
	routes.RegisterEventoRoutes(r, db)
	routes.RegisterIngressoRoutes(r, db)
	routes.RegisterCompraRoutes(r, db)
	routes.RegisterAvaliacaoRoutes(r, db)

	r.Run(":3000")
}
