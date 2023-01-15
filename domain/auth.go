package domain

import (
	"context"
	"github.com/NextTourPlan/domain/dto"
	"time"
)

type SignUpInput struct {
	ID              uint      `json:"id"`
	FullName        string    `json:"full_name"`
	Email           string    `json:"email"`
	Contact         string    `json:"contact"`
	Address         string    `json:"address"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Rating          int       `json:"rating"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SignInInput struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRepository interface {
	PostSignUP(ctx context.Context, ctr *SignUpInput) string
	PostSignIn(ctx context.Context, ctr *SignInInput) (*dto.JWTToken, error)
}

type AuthUseCase interface {
	PostSignUp(ctx context.Context, ctr *SignUpInput) string
	PostSignIn(ctx context.Context, ctr *SignInInput) (*dto.JWTToken, error)
}
