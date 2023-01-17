package usecase

import (
	"context"
	"fmt"
	"github.com/NextTourPlan/domain"
)

// New return new usecase for user
func New(repo domain.PlanForTourRepository) domain.PlanForTourUseCase {
	return &PlanForTourUseCase{
		repo: repo,
	}
}

type PlanForTourUseCase struct {
	repo domain.PlanForTourRepository
}

func (c *PlanForTourUseCase) Post(ctx context.Context, ctr *domain.PlanForTour) error {
	fmt.Println(ctx, ctr)
	return c.repo.Post(ctx, ctr)
}

func (c *PlanForTourUseCase) List(ctx context.Context, ctr *domain.PlanForTourCriteria) ([]*domain.PlanForTour, error) {
	return c.repo.List(ctx, ctr)
}

func (c *PlanForTourUseCase) Get(ctx context.Context, ctr *domain.PlanForTourCriteria) (*domain.TourDetails, error) {
	return c.repo.Get(ctx, ctr)
}

func (c *PlanForTourUseCase) Spots(ctx context.Context, ctr *domain.TourSpots) error {
	fmt.Println(ctx, ctr)
	return c.repo.Spots(ctx, ctr)
}

func (c *PlanForTourUseCase) SpotsList(ctx context.Context, ctr *domain.TourSpotsCriteria) ([]*domain.TourSpots, error) {
	return c.repo.SpotsList(ctx, ctr)
}

func (c *PlanForTourUseCase) Meals(ctx context.Context, ctr *domain.Meals) error {
	fmt.Println(ctx, ctr)
	return c.repo.Meals(ctx, ctr)
}

func (c *PlanForTourUseCase) MealsList(ctx context.Context, ctr *domain.MealsCriteria) ([]*domain.Meals, error) {
	return c.repo.MealsList(ctx, ctr)
}
