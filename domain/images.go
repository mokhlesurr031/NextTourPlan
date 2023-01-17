package domain

import "time"

type ImagesUploader struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ImgURLLocal  string    `json:"img_url_local"`
	ImgURLRemote string    `json:"img_url_remote"`
	DomainID     uint      `json:"domain_id"`
	TourID       uint      `json:"tour_id"`
	SpotID       uint      `json:"spot_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
