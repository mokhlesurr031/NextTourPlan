package domain

import (
	"context"
	"time"
)

type TourSpots struct {
	ID     uint   `json:"id"`
	TourID uint   `json:"tour_id"`
	Name   string `json:"name"`
	//Images      []*multipart.FileHeader `json:"images"`
	Description string `json:"description"`
}

type TourSpotsCriteria struct {
	ID     *uint   `json:"id"`
	TourID *uint   `json:"tour_id"`
	Name   *string `json:"name"`
	//Images      []*multipart.FileHeader `json:"images"`
	Description *string `json:"description"`
}

type Meals struct {
	ID        uint   `json:"id"`
	TourID    uint   `json:"tour_id"`
	BreakFast string `json:"break_fast"`
	Lunch     string `json:"lunch"`
	Dinner    string `json:"dinner"`
	Snacks    string `json:"snacks"`
	Others    string `json:"others"`
}

type MealsCriteria struct {
	ID        *uint   `json:"id"`
	TourID    *uint   `json:"tour_id"`
	BreakFast *string `json:"break_fast"`
	Lunch     *string `json:"lunch"`
	Dinner    *string `json:"dinner"`
	Snacks    *string `json:"snacks"`
	Others    *string `json:"others"`
}

type PlanForTour struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	PickUpLocation string    `json:"pick_up_location"`
	Description    string    `json:"description"`
	DayCount       string    `json:"day_count"`
	StartingAt     string    `json:"starting_at"`
	CostPerHead    string    `json:"cost_per_head"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PlanForTourCriteria struct {
	ID             *uint      `json:"id"`
	Name           *string    `json:"name"`
	PickUpLocation *string    `json:"pick_up_location"`
	Description    *string    `json:"description"`
	DayCount       *string    `json:"day_count"`
	StartingAt     *string    `json:"starting_at"`
	CostPerHead    *string    `json:"cost_per_head"`
	CreatedBy      *string    `json:"created_by"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type PlanForTourRepository interface {
	Post(ctx context.Context, ctr *PlanForTour) error
	List(ctx context.Context, ctr *PlanForTourCriteria) ([]*PlanForTour, error)

	Spots(ctx context.Context, ctr *TourSpots) error
	SpotsList(ctx context.Context, ctr *TourSpotsCriteria) ([]*TourSpots, error)

	Meals(ctx context.Context, ctr *Meals) error
	MealsList(ctx context.Context, ctr *MealsCriteria) ([]*Meals, error)
}

type PlanForTourUseCase interface {
	Post(ctx context.Context, ctr *PlanForTour) error
	List(ctx context.Context, ctr *PlanForTourCriteria) ([]*PlanForTour, error)

	Spots(ctx context.Context, ctr *TourSpots) error
	SpotsList(ctx context.Context, ctr *TourSpotsCriteria) ([]*TourSpots, error)

	Meals(ctx context.Context, ctr *Meals) error
	MealsList(ctx context.Context, ctr *MealsCriteria) ([]*Meals, error)
}
