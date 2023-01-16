package uploader

import (
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func NewHTTPHandler(r *chi.Mux) {
	r.Route("/api/uploader", func(r chi.Router) {
		r.Post("/", uploadHandler)
	})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	}
	files := r.MultipartForm.File["images"]
	for _, f := range files {
		file, _ := f.Open()
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		out, _ := os.Create("./img/upload/" + f.Filename)
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
	}
	_, err = fmt.Fprintf(w, "Upload complete!")
	if err != nil {
		return
	}
}
