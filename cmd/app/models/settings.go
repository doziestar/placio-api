package models

import (
	"gorm.io/gorm"
	"time"
)

type GeneralSettings struct {
	//gorm.Model
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	ID        string     `gorm:"primaryKey"`
	Language  string
	Theme     string
}

// GetGeneralSettings /*
func (g *GeneralSettings) GetGeneralSettings() (*GeneralSettings, error) {
	result := db.Where("id = ?", g.ID).First(&g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

// UpdateGeneralSettings /*
func (g *GeneralSettings) UpdateGeneralSettings() (*GeneralSettings, error) {
	result := db.Save(&g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

// CreateGeneralSettings /*
func (g *GeneralSettings) CreateGeneralSettings() (*GeneralSettings, error) {
	result := db.Create(&g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

// DeleteGeneralSettings /*
func (g *GeneralSettings) DeleteGeneralSettings() error {
	result := db.Delete(&g)
	return result.Error
}

// BeforeCreate OnCreateGeneralSettings /*
func (g *GeneralSettings) BeforeCreate(tx *gorm.DB) error {
	g.ID = GenerateID()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}
