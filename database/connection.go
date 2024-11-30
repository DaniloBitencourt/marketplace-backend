package database

import (
	"marketplace-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Função para inicializar
func InitDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/ticketmarket?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, db.AutoMigrate(
		&models.Usuario{}, &models.Evento{}, &models.Ingresso{}, &models.Compra{}, &models.Avaliacao{},
	)

}
