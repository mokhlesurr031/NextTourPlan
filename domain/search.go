package domain

import (
	"time"
)

type SearchTour struct {
	From        string    `json:"from"`
	To          string    `json:"to"`
	JourneyDate time.Time `json:"journey_date"`
}

//
//type SearchRepository interface {
//	SearchTour(ctx context.Context, ctr *SearchTour) error
//}
//
//type SearchUseCase interface {
//	SearchTour(ctx context.Context, ctr *SearchTour) error
//}
