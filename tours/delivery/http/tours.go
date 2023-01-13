package http

import (
	"encoding/json"
	"github.com/NextTourPlan/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type PlanForTourHandler struct {
	PlanForTourUseCase domain.PlanForTourUseCase
}

func NewHTTPHandler(r *chi.Mux, planForTourUseCase domain.PlanForTourUseCase) {
	handler := &PlanForTourHandler{
		PlanForTourUseCase: planForTourUseCase,
	}
	r.Route("/tours", func(r chi.Router) {
		r.Post("/post", handler.Post)
	})
}

type ReqPlanForTour struct {
	domain.PlanForTour
}

func (t *PlanForTourHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	req := ReqPlanForTour{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}
	ctx := r.Context()
	tour := domain.PlanForTour(req.PlanForTour)
	if err := t.PlanForTourUseCase.Post(ctx, &tour); err != nil {
		log.Println(err)
	}
}
