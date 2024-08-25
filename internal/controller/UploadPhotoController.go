package controller

import (
	"net/http"
	"payment/internal/service/profil"

	"github.com/gin-gonic/gin"
)

func UploadPhotoController(c *gin.Context) {
	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	defer file.Close()

	filename := header.Filename
	tempFilePath, err := profil.UploadPhotoService(file, filename, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	finalFilePath, err := profil.CompleteUploadService(tempFilePath, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully", "path": finalFilePath})
}
