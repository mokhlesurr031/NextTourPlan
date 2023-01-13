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
