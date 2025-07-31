package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var allowExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var allowMimeType = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

const maxSize = 5 << 20

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	//check extension in filename
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowExt[ext] {
		return "", errors.New("unsupported file extension")
	}

	//check size
	if fileHeader.Size > maxSize {
		return "", errors.New("file too large (Max 5 MB)")
	}

	//check filetye
	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", errors.New("failed to read file")
	}

	mimeType := http.DetectContentType(buffer)
	if !allowMimeType[mimeType] {
		return "", fmt.Errorf("unsupported file type: %s", mimeType)
	}

	//change file name abc.jpg
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	//create folder if not exist
	if err := os.MkdirAll("./upload", os.ModePerm); err != nil {
		return "", errors.New("failed to create directory")
	}

	//uploadDir "./upload" + filename "abc.jpg"
	savePath := filepath.Join(uploadDir, filename)
	if err := saveFile(fileHeader, savePath); err != nil {
		return "", err
	}

	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	// mở file hiện tại
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// tạo file trống
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// di chuyển file hiện tại vào file trống vừa tạo
	_, err = io.Copy(out, src)

	return err
}
