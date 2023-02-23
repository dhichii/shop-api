package helper

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckFileType(file *multipart.FileHeader) (string, error) {
	fileType := file.Header["Content-Type"][0]
	if fileType != "image/png" && fileType != "image/jpeg" {
		return "", errors.New("only allowed png or jpeg")
	}

	return strings.Split(fileType, "/")[1], nil
}

func ProcessImage(c *fiber.Ctx, file *multipart.FileHeader, feat string) (string, error) {
	fileType, err := CheckFileType(file)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("/%s.%s", uuid.New(), fileType)
	return fileName, c.SaveFile(file, fmt.Sprintf("./public/%s%s", feat, fileName))
}
