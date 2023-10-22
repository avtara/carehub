package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Base64ToRawBytes(b64Data string) (rawData []byte, err error) {
	rawData, err = base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		return
	}
	return
}

func SaveImageToLocalDir(b64Data, group string) (fullPath string, err error) {
	var rawData []byte
	rawData, err = Base64ToRawBytes(b64Data)
	if err != nil {
		err = fmt.Errorf("failed Base64ToRawBytes: %s", err.Error())
		return
	}

	if len(rawData) > models.MaxSize {
		err = errors.New("file size is larger than the limit")
		return
	}

	folderPath := "assets"
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			err = fmt.Errorf("file size is larger than the limit: %s", err.Error())
			return
		}
	}

	img, format, err := image.Decode(bytes.NewReader(rawData))
	if err != nil {
		err = fmt.Errorf("failed decode image: %s", err.Error())
		return
	}

	fileName := fmt.Sprintf("%s_%d.%s", group, time.Now().Unix(), format) // you should assign a proper filename, with an extension that matches the file type
	fullPath = filepath.Join(folderPath, fileName)
	outputFile, err := os.Create(fullPath)
	if err != nil {
		err = fmt.Errorf("failed creating file: %s", err.Error())
		return
	}
	defer outputFile.Close()

	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(outputFile, img, nil)
	case "png":
		err = png.Encode(outputFile, img)
	default:
		err = errors.New("unsupported image format")
		return
	}

	if err != nil {
		err = fmt.Errorf("error in saving image: %s", err.Error())
		return
	}

	return
}
