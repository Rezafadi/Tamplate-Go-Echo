package controllers

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"project-name/config"
	"strings"

	"github.com/labstack/echo/v4"
)

// UploadFile godoc
// @Summary Upload File
// @Description Upload File
// @Tags File
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "File to upload (PDF, JPEG, JPG, PNG)"
// @Success 200
// @Router /v1/file/upload [post]
// @Security JwtToken
func UploadFile(c echo.Context) error {
	// Read form file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Check file type
	allowedTypes := map[string]bool{
		"application/pdf": true,
		"image/jpeg":      true,
		"image/jpg":       true,
		"image/png":       true,
	}

	// Get file header to check MIME type
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return err
	}
	src.Seek(0, io.SeekStart) // Reset the read pointer to the start of the file

	mimeType := http.DetectContentType(buffer)
	if !allowedTypes[mimeType] {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  400,
			"message": "File type not allowed. Only PDF, JPEG, JPG, and PNG are accepted.",
		})
	}

	// Create directory path for uploads if not exists
	uploadDir := config.LoadConfig().DirPath
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	cleanFilename := strings.ReplaceAll(file.Filename, " ", "")

	randomString, _ := GenerateRandomString(10)

	filename := randomString + "_" + cleanFilename

	// Destination
	dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDir, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "Upload Success",
		"data":    filename,
	})
}

// UploadMultipleFiles godoc
// @Summary Upload Multiple Files
// @Description Upload Multiple Files
// @Tags File
// @Accept multipart/form-data
// @Produce application/json
// @Param files formData file true "Files to upload (PDF, JPEG, JPG, PNG)"
// @Success 200
// @Router /v1/file/upload-multiple [post]
// @Security JwtToken
func UploadMultipleFiles(c echo.Context) error {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  400,
			"message": "Invalid form data",
		})
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  400,
			"message": "No files to upload",
		})
	}

	// Allowed file types
	allowedTypes := map[string]bool{
		"application/pdf": true,
		"image/jpeg":      true,
		"image/jpg":       true,
		"image/png":       true,
	}

	uploadDir := config.LoadConfig().DirPath
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	uploadedFiles := []string{}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Check file type
		buffer := make([]byte, 512)
		_, err = src.Read(buffer)
		if err != nil {
			return err
		}
		src.Seek(0, io.SeekStart)

		mimeType := http.DetectContentType(buffer)
		if !allowedTypes[mimeType] {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  400,
				"message": fmt.Sprintf("File %s has an invalid type", file.Filename),
			})
		}

		cleanFilename := strings.ReplaceAll(file.Filename, " ", "")
		randomString, _ := GenerateRandomString(10)
		filename := randomString + "_" + cleanFilename

		dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDir, filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy file to destination
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		uploadedFiles = append(uploadedFiles, filename)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "Upload Success",
		"data":    uploadedFiles,
	})
}

func GenerateRandomString(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result[i] = chars[randomIndex.Int64()]
	}
	return string(result), nil
}

func DeleteFile(filename string) error {

	e := os.Remove(fmt.Sprintf(config.LoadConfig().DirPath+"%s", filename))
	if e != nil {
		fmt.Println("Failed to delete file. Error:", e)
	}

	fmt.Println("Success to delete file.")

	return nil
}
