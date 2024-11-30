package models

import "time"

type Usuario struct {
	ID       uint      `gorm:"primaryKey"`
	Nome     string    `gorm:"size:100;not null"`
	Email    string    `gorm:"size:150;unique;not null"`
	Senha    string    `gorm:"size:255;not null"`
	Telefone string    `gorm:"size:15"`
	Tipo     string    `gorm:"type:enum('comprador','organizador');not null"`
	Pontos   int       `gorm:"default:0"`
	CriadoEm time.Time `gorm:"autoCreateTime"`
}
