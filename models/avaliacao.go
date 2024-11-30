package models

import "time"

type Avaliacao struct {
	ID         uint      `gorm:"primaryKey"`
	UsuarioID  uint      `gorm:"not null"`
	EventoID   uint      `gorm:"not null"`
	Nota       int       `gorm:"not null"`
	Comentario string    `gorm:"type:text"`
	CriadoEm   time.Time `gorm:"autoCreateTime"`
}
