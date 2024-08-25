package profil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"payment/internal/entity"
	"payment/internal/helper"
	"payment/internal/repository"
	"payment/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadPhotoService(file io.Reader, filename string, c *gin.Context) (string, error) {
	token, err := helper.GetToken(c)
	if err != nil {
		return "", err
	}
	tempDir := "./upload/temp"
	os.MkdirAll(tempDir, os.ModePerm)
	tempFilename := fmt.Sprintf("%d_%s", token.Id, filename)
	tempFilePath := filepath.Join(tempDir, tempFilename)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", err
	}
	chunkSize := 1024 * 1024
	fileInfo, err := tempFile.Stat()
	if err != nil {
		return "", err
	}

	totalChunks := int(fileInfo.Size()/int64(chunkSize)) + 1
	tempFile.Seek(0, 0)
	for i := 0; i < totalChunks; i++ {
		partSize := int64(chunkSize)
		if i == totalChunks-1 {
			partSize = fileInfo.Size() - int64(i)*int64(chunkSize)
		}

		partBuffer := make([]byte, partSize)
		_, err := tempFile.Read(partBuffer)
		if err != nil {
			return "", err
		}

		partFileName := fmt.Sprintf("%s.part%d", tempFilename, i+1)
		partFilePath := filepath.Join(tempDir, partFileName)

		partFile, err := os.Create(partFilePath)
		if err != nil {
			return "", err
		}
		defer partFile.Close()

		partFile.Write(partBuffer)
	}

	return tempFilePath, nil
}

func CompleteUploadService(tempFilePath string, c *gin.Context) (string, error) {
	token, err := helper.GetToken(c)
	if err != nil {
		return "", err
	}
	user := entity.User{}
	var finalFilePath string
	err = config.GetPgsqlDB().Transaction(func(tx *gorm.DB) error {

		err := repository.GetUserByEmail(&user, token.Email, tx, c)
		if err != nil {
			return err
		}
		uploadDir := "./uploads/photos"
		os.MkdirAll(uploadDir, os.ModePerm)
		finalFilename := filepath.Base(tempFilePath)
		finalFilePath = filepath.Join(uploadDir, finalFilename)
		err = os.Rename(tempFilePath, finalFilePath)
		if err != nil {
			return err
		}

		user.Photo = finalFilePath
		err = repository.UpdateUserById(&user, user.ID, tx, c)
		if err != nil {
			return err
		}

		return nil
	})

	return finalFilePath, nil
}
