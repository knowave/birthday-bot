package service

import (
	"time"

	"birthday-bot/internal/domain"
	"birthday-bot/internal/repository"
)

type SlackUserService struct {
    slackUserRepo *repository.SlackUserRepository
}

func NewSlackUserService(slackUserRepo *repository.SlackUserRepository) *SlackUserService {
    return &SlackUserService{slackUserRepo: slackUserRepo}
}

func (s *SlackUserService) GetByID(id uint) (*domain.SlackUser, error) {
    return s.slackUserRepo.FindByID(id)
}

func (s *SlackUserService) GetByEmail(email string) (*domain.SlackUser, error) {
    return s.slackUserRepo.FindByEmail(email)
}

func (s *SlackUserService) GetTodayBirthdays() ([]domain.SlackUser, error) {
    now := time.Now()
    return s.slackUserRepo.FindByBirthday(now.Month(), now.Day())
}

func (s *SlackUserService) GetAll() ([]domain.SlackUser, error) {
    return s.slackUserRepo.FindAll()
}

func (s *SlackUserService) Create(name, email, slackUserID string, birthday time.Time) (*domain.SlackUser, error) {
    user := domain.NewSlackUser(name, email, slackUserID, birthday)
    err := s.slackUserRepo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (s *SlackUserService) Update(id uint, name, email string, birthday time.Time) (*domain.SlackUser, error) {
    user, err := s.slackUserRepo.FindByID(id)
    if err != nil {
        return nil, err
    }

    user.Name = name
    user.UpdateEmail(email)
    user.UpdateBirthday(birthday)

    err = s.slackUserRepo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (s *SlackUserService) Delete(id uint) error {
    return s.slackUserRepo.Delete(id)
}