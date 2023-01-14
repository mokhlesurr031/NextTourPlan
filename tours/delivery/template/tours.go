package template

import (
	"fmt"
	"github.com/NextTourPlan/domain"
	"github.com/NextTourPlan/internal/conn"
	"github.com/go-chi/chi"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func NewHTTPHandler(r *chi.Mux) {

	r.Route("/tours", func(r chi.Router) {
		r.Get("/", Post)
		r.Post("/", Post)
		r.Get("/list", Get)
	})
}

func Post(w http.ResponseWriter, r *http.Request) {
	db := conn.DefaultDB()
	//closeDB, _ := db.DB()
	//defer func(closeDB *sql.DB) {
	//	err := closeDB.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}(closeDB)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/tours/create.html")
		err := t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		dayCount, _ := strconv.Atoi(r.Form["day_count"][0])
		costPerHead, _ := strconv.ParseFloat(r.Form["cost_per_head"][0], 64)
		createdBy, _ := strconv.ParseUint(r.Form["created_by"][0], 10, 64)

		req := domain.PlanForTour{
			Name:           r.Form["name"][0],
			Description:    r.Form["description"][0],
			PickUpLocation: r.Form["pick_up_location"][0],
			DayCount:       dayCount,
			CostPerHead:    costPerHead,
			CreatedBy:      uint(createdBy),
		}

		if err := db.Create(&req).Error; err != nil {
			log.Println(err)
		}
		fmt.Println(req)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	db := conn.DefaultDB()
	//closeDB, _ := db.DB()
	//defer func(closeDB *sql.DB) {
	//	err := closeDB.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}(closeDB)
	if r.Method == "GET" {
		req := []domain.PlanForTour{}
		if err := db.Find(&req).Error; err != nil {
			log.Println(err)
		}

		t, _ := template.ParseFiles("template/tours/list.html")
		err := t.Execute(w, req)
		if err != nil {
			log.Println(err)
		}
	}
}
