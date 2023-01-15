package usecase

import (
	"context"
	"github.com/NextTourPlan/domain"
	"github.com/NextTourPlan/domain/dto"
)

// New return new usecase for user
func New(repo domain.AuthRepository) domain.AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

type AuthUseCase struct {
	repo domain.AuthRepository
}

func (a *AuthUseCase) PostSignUp(ctx context.Context, ctr *domain.SignUpInput) string {
	return a.repo.PostSignUP(ctx, ctr)
}

func (a *AuthUseCase) PostSignIn(ctx context.Context, ctr *domain.SignInInput) (*dto.JWTToken, error) {
	return a.repo.PostSignIn(ctx, ctr)
}
