package repository

import (
	"birthday-bot/internal/domain"

	"gorm.io/gorm"
)

type ConfigRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	return &ConfigRepository{ db: db }
}

func (r *ConfigRepository) FindByKey(key string) (*domain.Config, error) {
	var config domain.Config
	err := r.db.Where("key = ?", key).First(&config).Error

	if err != nil {
		return nil, err
	}

	return &config, nil	
}

func (r *ConfigRepository) Save(config *domain.Config) error {
	return r.db.Save(config).Error
}