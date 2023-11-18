package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

func (env *Env) HandleRetrieveImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		table := params["table"]
		id := params["id"]
		url := "./images/" + table + "/" + id

		http.ServeFile(w, r, url)
	}
}

func (env *Env) HandleImageUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		table := params["table"]
		//handle multipart form part
		formFile, fileHeader, err := r.FormFile("file")
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		defer formFile.Close()

		resultPath, err := env.db.UploadImage(formFile, fileHeader, table)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		env.encodeResponse(w, &ResponseModel{Message: "success", Data: resultPath})
	}
}

func (env *Env) HandleDeleteImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		img := models.ImageSearch{}
		if ok := env.decodeRequest(w, r, &img); !ok {
			return
		}
		err := env.db.DeleteImage(img.ImageURL)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: "Image removed."})

	}
}
