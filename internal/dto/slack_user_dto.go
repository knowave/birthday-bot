package dto

import "time"


type CreateSlackUserRequest struct {
    Name        string		`json:"name"`
    Email       string		`json:"email"`
    SlackUserID string		`json:"slack_user_id"`
    Birthday    time.Time	`json:"birthday"`
}