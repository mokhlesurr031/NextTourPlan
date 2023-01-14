package http

import (
	"encoding/json"
	"github.com/NextTourPlan/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
}

func NewHTTPHandler(r *chi.Mux, authUseCase domain.AuthUseCase) {
	handler := &AuthHandler{
		AuthUseCase: authUseCase,
	}
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/", handler.PostSignUP)
		r.Post("/login", handler.PostSignIn)
	})
}

type ReqSignUp struct {
	domain.SignUpInput
}

type ReqSignIn struct {
	domain.SignInInput
}

func (a *AuthHandler) PostSignUP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ReqSignUp{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	signup := domain.SignUpInput(req.SignUpInput)
	if err := a.AuthUseCase.PostSignUp(ctx, &signup); err != nil {
		log.Println(err)
	}
}

func (a *AuthHandler) PostSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	req := ReqSignIn{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	signIn := domain.SignInInput(req.SignInInput)
	resp, err := a.AuthUseCase.PostSignIn(ctx, &signIn)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(er)
	}
}
