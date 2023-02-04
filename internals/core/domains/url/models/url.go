package models

import "time"

type Url struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	MainUrl   string    `gorm:"type:varchar(199);not null" json:"main_url,omitempty"`
	ShortUrl  string    `gorm:"type:varchar(199);not null" json:"short_url,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type UrlStoreRequest struct {
	MainUrl   string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
