package domain

import "time"

type ImagesUploader struct {
	ID        uint      `json:"id"`
	ImgPath   string    `json:"img_path"`
	DomainID  uint      `json:"domain_id"`
	TourID    uint      `json:"tour_id"`
	SpotID    uint      `json:"spot_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
