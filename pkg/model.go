package pkg

import (
	"time"
)

type Model struct {
	ID        string    `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()"`
}
