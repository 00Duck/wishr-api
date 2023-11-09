package app

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/00Duck/wishr-api/models"
	"github.com/google/uuid"
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

		//open file from fileheader above
		file, err := fileHeader.Open()
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		defer file.Close()

		//make a buffer to read the first 512 bytes of the file
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		//check the content type and only allow certain types
		fileType := http.DetectContentType(buff)
		if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
			env.encodeResponse(w, &ResponseModel{Message: "Error: invalid image type. Please use only .png, .jpg, or .jpeg"})
			return
		}

		//start at the beginning of the file again
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		//ensure the correct directory structure exists on the server
		imagePath := "./images/" + table
		err = os.MkdirAll(imagePath, os.ModePerm)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		//TODO: auto resize file

		//generate random file name
		fileName := uuid.New().String()
		fileName = strings.Replace(fileName, "-", "", -1)

		//create the shell of the new file on the server at the correct location
		fullFileName := fileName + filepath.Ext(fileHeader.Filename)
		f, err := os.Create(imagePath + "/" + fullFileName)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		defer f.Close()

		//copy the contents of the uploaded file into our new location
		_, err = io.Copy(f, file)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}

		env.encodeResponse(w, &ResponseModel{Message: "success", Data: "/images/" + table + "/" + fullFileName})
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
