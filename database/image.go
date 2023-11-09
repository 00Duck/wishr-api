package database

import (
	"errors"
	"os"
	"strings"
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
