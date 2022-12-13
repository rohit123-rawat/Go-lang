package domain

import "time"

type Testing struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	testCount    uint64    `gorm:"not null;" json:"test_count"`
}
