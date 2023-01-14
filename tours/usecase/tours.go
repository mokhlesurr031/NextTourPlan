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
