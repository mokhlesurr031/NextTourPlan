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

		r.Post("/spots", handler.Spots)
		r.Get("/spots/list", handler.SpotsList)

		r.Post("/meals", handler.Meals)
		r.Get("/meals/list", handler.MealsList)

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

	err := json.NewEncoder(w).Encode(tour)
	if err != nil {
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

type ReqTourSpots struct {
	domain.TourSpots
}

func (t *PlanForTourHandler) Spots(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	req := ReqTourSpots{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}
	ctx := r.Context()
	spots := domain.TourSpots(req.TourSpots)
	if err := t.PlanForTourUseCase.Spots(ctx, &spots); err != nil {
		log.Println(err)
	}
}

func (t *PlanForTourHandler) SpotsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx := r.Context()
	spots := &domain.TourSpotsCriteria{}
	spotsList, err := t.PlanForTourUseCase.SpotsList(ctx, spots)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(spotsList)
	if err != nil {
		log.Println(er)
	}
}

type ReqMeals struct {
	domain.Meals
}

func (t *PlanForTourHandler) Meals(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	req := ReqMeals{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}
	ctx := r.Context()
	meals := domain.Meals(req.Meals)
	if err := t.PlanForTourUseCase.Meals(ctx, &meals); err != nil {
		log.Println(err)
	}
}

func (t *PlanForTourHandler) MealsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx := r.Context()
	meals := &domain.MealsCriteria{}
	mealsList, err := t.PlanForTourUseCase.MealsList(ctx, meals)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(mealsList)
	if err != nil {
		log.Println(er)
	}
}
