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

func storeLocal(file multipart.File, filePath string, fileName string) error {
	// Create a new file to store the image
	f, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return nil
}

func uploadSpotImgHandler(w http.ResponseWriter, r *http.Request) {
	db := conn.DefaultDB()
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}
	files, _, _ := r.FormFile("images")

	userID := r.FormValue("user_id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}
	domainID := r.FormValue("domain_id")
	domainIDUint, err := strconv.ParseUint(domainID, 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}
	spotID := r.FormValue("spot_id")
	spotIDUint, err := strconv.ParseUint(spotID, 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}
	tourID := r.FormValue("tour_id")
	tourIDUint, err := strconv.ParseUint(tourID, 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}

	filePath := "./img/upload/tours/"
	fileName := "tours_uid" + userID + "_did" + domainID + "_tid" + tourID + "_sid" + spotID + ".jpg"

	er := storeLocal(files, filePath, fileName)
	if er != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}

	data := &domain.ImagesUploader{DomainID: uint(domainIDUint), ImgPath: filePath + fileName, SpotID: uint(spotIDUint), TourID: uint(tourIDUint), UserID: uint(userIDUint)}
	db.Create(&data)

	_, _ = fmt.Fprintf(w, "Upload complete!")

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
