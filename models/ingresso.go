package models

import "time"

type Ingresso struct {
	ID         uint      `gorm:"primaryKey"`
	EventoID   uint      `gorm:"not null"`
	Preco      float64   `gorm:"not null"`
	Tipo       string    `gorm:"size:50;not null"`
	CriadoEm   time.Time `gorm:"autoCreateTime"`
	Disponivel bool      `gorm:"default:true"`
}
