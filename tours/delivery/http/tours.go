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
	r.Route("/api/tours", func(r chi.Router) {
		r.Post("/post", handler.Post)
		r.Get("/list", handler.List)
		//r.Post("/list", handler.SearchTour)
	})
}

type ReqPlanForTour struct {
	domain.PlanForTour
}

type ReqSearchTour struct {
	domain.SearchTour
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

//func (t *PlanForTourHandler) SearchTour(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("content-type", "application/json")
//	req := ReqSearchTour{}
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		log.Println(err)
//	}
//	ctx := r.Context()
//	search := domain.SearchTour(req.SearchTour)
//	if err := t.PlanForTourUseCase.Post(ctx, &search); err != nil {
//		log.Println(err)
//	}
//}

func (t *PlanForTourHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx := r.Context()
	tours := &domain.PlanForTourCriteria{}
	toursList, err := t.PlanForTourUseCase.List(ctx, tours)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(toursList)
	if err != nil {
		log.Println(er)
	}
}
