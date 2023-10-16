package utils

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func getCloudinary() *cloudinary.Cloudinary {
	cd, err := cloudinary.New()

	if err != nil {
		panic(err)
	}

	return cd
}

var GetCloudinary = sync.OnceValue(getCloudinary)

func UploadFile(c *gin.Context, filetype string) (url string, ok bool) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File missing",
		})
		return "", false
	}

	uploadedFile, err := file.Open()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File cannot be parsed",
		})
		return "", false
	}

	defer uploadedFile.Close()

	buffer := make([]byte, 512)

	_, err = uploadedFile.Read(buffer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File cannot be parsed",
		})
		return "", false
	}

	mimeType := http.DetectContentType(buffer)

	if !strings.Contains(mimeType, filetype) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s only", filetype),
		})
		return
	}

	// Not secure
	ext := filepath.Ext(file.Filename)
	publicID := fmt.Sprintf("%v-%v", time.Now().UnixNano(), rand.Intn(999_999_999))
	path := fmt.Sprintf("tmp/files/%s.%s", publicID, ext)

	c.SaveUploadedFile(file, path)

	result, err := GetCloudinary().Upload.Upload(context.Background(), path, uploader.UploadParams{
		PublicID:     publicID,
		Folder:       "/reddit/",
		Type:         api.Upload,
		ResourceType: "auto",
	})

	if err != nil {
		panic(err)
	}

	return result.SecureURL, true
}
