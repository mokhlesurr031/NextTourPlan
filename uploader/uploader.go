package uploader

import (
	"fmt"
	"github.com/NextTourPlan/domain"
	"github.com/NextTourPlan/internal/conn"
	"github.com/go-chi/chi"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func NewHTTPHandler(r *chi.Mux) {
	r.Route("/api/uploader", func(r chi.Router) {
		r.Post("/", uploadSpotImgHandler)
		r.Get("/img", getSpotImgHandler)
	})
}

type ImgData struct {
	DomainID uint
	ImgURl   string
}

func uploadSpotImgHandler(w http.ResponseWriter, r *http.Request) {
	db := conn.DefaultDB()
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	}
	files := r.MultipartForm.File["images"]
	domainID := r.FormValue("domain_id")
	domainIDUint, err := strconv.ParseUint(domainID, 10, 64)
	if err != nil {
		return
	}

	spotID := r.FormValue("spot_id")
	spotIDUint, err := strconv.ParseUint(spotID, 10, 64)
	if err != nil {
		return
	}

	tourID := r.FormValue("tour_id")
	tourIDUint, err := strconv.ParseUint(tourID, 10, 64)
	if err != nil {
		return
	}

	for _, f := range files {
		file, _ := f.Open()
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		out, _ := os.Create("./img/upload/tours/tours_" + f.Filename)

		fmt.Println("FileName", f.Filename)

		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				return
			}
		}(out)
		_, err := io.Copy(out, file)
		if err != nil {
			return
		}
		data := &domain.ImagesUploader{DomainID: uint(domainIDUint), ImgPath: "tours_" + f.Filename, SpotID: uint(spotIDUint), TourID: uint(tourIDUint)}
		db.Create(&data)
	}

	_, err = fmt.Fprintf(w, "Upload complete!")
	if err != nil {
		return
	}
}

func getSpotImgHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//imageName := vars["imageName"]
	file, _ := os.Open("./img/upload/tours/" + "tours_271713555_3182650491970063_2216618754507835589_n.jpg")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	_, err := io.Copy(w, file)
	if err != nil {
		return
	}
}
