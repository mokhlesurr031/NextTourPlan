package repository

import (
	"context"
	"github.com/NextTourPlan/domain"
	"gorm.io/gorm"
)

// New return Category SQL(MySQL/PostgreSQL) storage implementation
func New(db *gorm.DB) domain.PlanForTourRepository {
	return &TourSqlStorage{
		db: db,
	}
}

// CategorySqlStorage return a SQL implementation of the storage
type TourSqlStorage struct {
	db *gorm.DB
}

func (c *TourSqlStorage) Post(ctx context.Context, tour *domain.PlanForTour) error {
	db := c.db
	db.Create(tour)
	return nil
}

func (c *TourSqlStorage) List(ctx context.Context, ctr *domain.PlanForTourCriteria) ([]*domain.PlanForTour, error) {
	qry := c.db

	toursList := make([]*domain.PlanForTour, 0)
	if err := qry.WithContext(ctx).Find(&toursList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toursList, nil
}

func (c *TourSqlStorage) Spots(ctx context.Context, spots *domain.TourSpots) error {
	db := c.db
	db.Create(spots)
	return nil
}

func (c *TourSqlStorage) SpotsList(ctx context.Context, ctr *domain.TourSpotsCriteria) ([]*domain.TourSpots, error) {
	qry := c.db

	spotsList := make([]*domain.TourSpots, 0)
	if err := qry.WithContext(ctx).Find(&spotsList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return spotsList, nil
}

func (c *TourSqlStorage) Meals(ctx context.Context, meals *domain.Meals) error {
	db := c.db
	db.Create(meals)
	return nil
}

func (c *TourSqlStorage) MealsList(ctx context.Context, ctr *domain.MealsCriteria) ([]*domain.Meals, error) {
	qry := c.db

	mealsList := make([]*domain.Meals, 0)
	if err := qry.WithContext(ctx).Find(&mealsList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mealsList, nil
}
