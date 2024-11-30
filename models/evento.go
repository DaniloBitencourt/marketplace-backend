package models

import "time"

type Evento struct {
	ID            uint      `gorm:"primaryKey"`
	Nome          string    `gorm:"size:150;not null"`
	Descricao     string    `gorm:"type:text;not null"`
	Local         string    `gorm:"size:255;not null"`
	DataEvento    time.Time `gorm:"type:datetime;not null"`
	OrganizadorID uint      `gorm:"not null"`
	CriadoEm      time.Time `gorm:"autoCreateTime"`
}
