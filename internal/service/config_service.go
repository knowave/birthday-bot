package service

import (
	"birthday-bot/internal/domain"
	"birthday-bot/internal/repository"
)

type ConfigService struct {
	repository *repository.ConfigRepository
}

func NewConfigService(repository *repository.ConfigRepository) *ConfigService {
	return &ConfigService{repository: repository}
}

func (s *ConfigService) GetValue(key string) (string ,error) {
	config, err := s.repository.FindByKey(key)

	if err != nil {
		return "", err
	}

	return config.Value, err
}

func (s *ConfigService) SetValue(key string, value string) error {
	config, err := s.repository.FindByKey(key)

	if err != nil {
		config = domain.NewConfig(key, value)
	} else {
		config.UpdateValue(value)
	}

	return s.repository.Save(config)
}

func (s *ConfigService) GetSlackBotToken() (string, error) {
	return s.GetValue("SLACK_BOT_TOKEN")
}

func (s *ConfigService) GetSlackChannel() (string, error) {
	return s.GetValue("SLACK_CHANNEL")
}