package database

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func (d *DB) DeleteImage(imgPath string) error {
	if imgPath == "" {
		return errors.New("Error: no image path to delete.")
	}

	if strings.Contains(imgPath, "..") || strings.Contains(imgPath, "\\\\") {
		return errors.New("Error: invalid image path to delete.")
	}

	imgPath = strings.ReplaceAll(imgPath, "..", "") //redundancy is key is key

	err := os.Remove("." + imgPath) //add a dot for current directory
	if err != nil {
		return errors.New("Error: invalid image path to delete.")
	}
	return nil
}

// UploadImage provide a formfile (from upload), its fileheadder, and the name of the table/folder to store the file.
// returns the file path of the stored file and any error
func (d *DB) UploadImage(formFile multipart.File, fileHeader *multipart.FileHeader, table string) (string, error) {
	//open file from fileheader above
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	//make a buffer to read the first 512 bytes of the file
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		return "", err
	}

	//check the content type and only allow certain types
	fileType := http.DetectContentType(buff)
	isHeic := detectHeic(fileType, buff)

	if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" && !isHeic {
		return "", errors.New("Error: invalid image type. Please use only .png, .jpg, .heic, or .jpeg")
	}

	//start at the beginning of the file again
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	//ensure the correct directory structure exists on the server
	imagePath := "./images/" + table
	err = os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		return "", err
	}

	//TODO: auto resize file

	//generate random file name
	fileName := uuid.New().String()
	fileName = strings.Replace(fileName, "-", "", -1)

	//create the shell of the new file on the server at the correct location
	fullFileName := fileName + filepath.Ext(fileHeader.Filename)
	f, err := os.Create(imagePath + "/" + fullFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	//copy the contents of the uploaded file into our new location
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	//convert HEIC format to JPEG
	if isHeic {
		imagick.Initialize()
		defer imagick.Terminate()
		_, err := imagick.ConvertImageCommand([]string{
			"convert", imagePath + "/" + fullFileName, imagePath + "/" + fileName + ".jpeg",
		})
		//remove old uploaded heic if it's there
		os.Remove(imagePath + "/" + fullFileName)
		if err != nil {
			return "", err
		}
		//reset full file name so it has the converted extension
		fullFileName = fileName + ".jpeg"
	}
	return "/images/" + table + "/" + fullFileName, nil
}

// Detection based on signature from https://en.wikipedia.org/wiki/List_of_file_signatures
func detectHeic(fileType string, buff []byte) bool {
	heicSignature := []byte("\x66\x74\x79\x70\x68\x65\x69\x63") //ftypheic
	// HTTP's DetectContentType does not detect HEIC and returns this more general type. If we fail this starting point,
	// then there is no need to detect further
	if fileType != "application/octet-stream" {
		return false
	}

	detectSlice := buff[4:12]
	return bytes.Equal(heicSignature[:], detectSlice[:])
}
