package repository

import (
	"birthday-bot/internal/domain"
	"time"

	"gorm.io/gorm"
)

type SlackUserRepository struct {
	db *gorm.DB
}

func NewSlackUserRepository(db *gorm.DB) *SlackUserRepository {
	return &SlackUserRepository{db: db}
}

func (r *SlackUserRepository) FindByID(id uint) (*domain.SlackUser, error) {
	var slackUser domain.SlackUser
	err := r.db.First(&slackUser, id).Error

	if err != nil {
		return nil, err
	}

	return &slackUser, err
}

func (r *SlackUserRepository) FindByEmail(email string) (*domain.SlackUser, error) {
	var slackUser domain.SlackUser
	err := r.db.Where("email = ?", email).First(&slackUser).Error

	if err != nil {
		return nil, err
	}

	return &slackUser, err
}

func (r *SlackUserRepository) FindByBirthday(month time.Month, day int) ([]domain.SlackUser, error) {
	var slackUsers []domain.SlackUser
	err := r.db.Where("MONTH(birthday) = ? AND DAY(birthday) = ?", month, day).Find(&slackUsers).Error

	if err != nil {
		return nil, err
	}

	return slackUsers, nil
}

func (r *SlackUserRepository) FindAll(sort, order string) ([]domain.SlackUser, error) {
	var users []domain.SlackUser
	err := r.db.Order(slackUserOrderClause(sort, order)).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func slackUserOrderClause(sort, order string) string {
	direction := "ASC"
	if order == "desc" {
		direction = "DESC"
	}

	switch sort {
	case "name":
		return "name " + direction
	default:
		return "MONTH(birthday) " + direction + ", DAY(birthday) " + direction + ", birthday " + direction
	}
}

func (r *SlackUserRepository) Save(user *domain.SlackUser) error {
	return r.db.Save(user).Error
}

func (r *SlackUserRepository) SaveAllInBatches(slackUsers []*domain.SlackUser, batchSize int) error {
	if len(slackUsers) == 0 {
		return nil
	}

	if batchSize <= 0 {
		batchSize = 100
	}

	return r.db.CreateInBatches(slackUsers, batchSize).Error
}

func (r *SlackUserRepository) Delete(id uint) error {
	return r.db.Delete(&domain.SlackUser{}, id).Error
}
