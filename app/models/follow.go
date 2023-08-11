package models

import "time"

type Follow struct {
	ID          string    `gorm:"primaryKey,unique"`
	FollowerID  string    `gorm:"column:follower_id"`
	FollowingID string    `gorm:"column:following_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
