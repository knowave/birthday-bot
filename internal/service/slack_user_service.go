package service

import (
	"time"

	"birthday-bot/internal/domain"
	"birthday-bot/internal/dto"
	"birthday-bot/internal/repository"
)

type SlackUserService struct {
    repository *repository.SlackUserRepository
}

func NewSlackUserService(repository *repository.SlackUserRepository) *SlackUserService {
    return &SlackUserService{repository: repository}
}

func (s *SlackUserService) GetByID(id uint) (*domain.SlackUser, error) {
    return s.repository.FindByID(id)
}

func (s *SlackUserService) GetByEmail(email string) (*domain.SlackUser, error) {
    return s.repository.FindByEmail(email)
}

func (s *SlackUserService) GetTodayBirthdays() ([]domain.SlackUser, error) {
    now := time.Now()
    return s.repository.FindByBirthday(now.Month(), now.Day())
}

func (s *SlackUserService) GetAll() ([]domain.SlackUser, error) {
    return s.repository.FindAll()
}

func (s *SlackUserService) Create(name, email, slackUserID string, birthday time.Time) (*domain.SlackUser, error) {
    user := domain.NewSlackUser(name, email, slackUserID, birthday)
    err := s.repository.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (s *SlackUserService) Update(id uint, name, email string, birthday time.Time) (*domain.SlackUser, error) {
    user, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }

    user.Name = name
    user.UpdateEmail(email)
    user.UpdateBirthday(birthday)

    err = s.repository.Save(user)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *SlackUserService) Delete(id uint) error {
    return s.repository.Delete(id)
}

func (s *SlackUserService) CreateAll(dto []dto.CreateSlackUserRequest) error {
	slackUsers := make([]*domain.SlackUser, len(dto))

	for i, item := range dto {
		slackUsers[i] = domain.NewSlackUser(item.Name, item.Email, item.SlackUserID, item.Birthday)
	}

	return s.repository.SaveAllInBatches(slackUsers, 100)
}