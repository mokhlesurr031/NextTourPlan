package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NextTourPlan/domain"
	"github.com/NextTourPlan/internal/conn"
	"github.com/go-chi/chi"
	"io"
	"log"
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

func storeLocally(file multipart.File, filePath string, fileName string) error {
	// Create a new file to store the image
	f, err := os.Create(filePath + fileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return nil
}

func storeRemote(file multipart.File, fileName string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("fileName", fileName)
	io.Copy(part, file)
	writer.Close()

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "https://api.upload.io/v2/accounts/12a1xvh/uploads/form_data", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer public_12a1xvh6D2kN8NT8fAZpvmG9sxvZ")

	// Send the request
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	type Response struct {
		Files []struct {
			FormDataFieldName string `json:"formDataFieldName"`
			AccountId         string `json:"accountId"`
			FilePath          string `json:"filePath"`
			FileUrl           string `json:"fileUrl"`
		} `json:"files"`
	}
	var r Response
	// Print the response
	json.NewDecoder(resp.Body).Decode(&r)

	imgUrl := r.Files[0].FileUrl

	return imgUrl
}

func uploadSpotImgHandler(w http.ResponseWriter, r *http.Request) {
	db := conn.DefaultDB()
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}
	files, handler, _ := r.FormFile("images")
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

	//fileNameExt := "tours_uid" + userID + "_did" + domainID + "_tid" + tourID + "_sid" + spotID
	fileName := handler.Filename
	filePath := "./img/upload/tours/"

	er := storeLocally(files, filePath, fileName)
	if er != nil {
		_, _ = fmt.Fprintf(w, "Upload failed!")
	}

	imgUrlRemote := storeRemote(files, fileName)

	data := &domain.ImagesUploader{DomainID: uint(domainIDUint), ImgURLRemote: imgUrlRemote, ImgURLLocal: filePath + fileName, SpotID: uint(spotIDUint), TourID: uint(tourIDUint), UserID: uint(userIDUint)}
	db.Create(&data)

	_, _ = fmt.Fprintf(w, "Upload complete!")

}

func getSpotImgHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./img/upload/tours/tours_uid1_did1_tid1_sid2.jpg")
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
