package models

import "time"

type Compra struct {
	ID         uint      `gorm:"primaryKey"`
	UsuarioID  uint      `gorm:"not null"`
	IngressoID uint      `gorm:"not null"`
	Quantidade int       `gorm:"not null"`
	PrecoTotal float64   `gorm:"not null"`
	Status     string    `gorm:"size:50;not null"`
	CriadoEm   time.Time `gorm:"autoCreateTime"`
}
