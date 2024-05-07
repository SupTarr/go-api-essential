package utils

import (
	"time"

	"gorm.io/gorm"
)

type status string

const (
	Success      status = "SUCCESS"
	Fail         status = "FAILED"
	DataNotFound status = "DATA_NOT_FOUND"
)

type Response struct {
	Status  status `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
