package handler

import (
	"io"
	"mulfinance/pkg/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Routes(group *echo.Group) {
	group.POST("/upload", handler.Upload)
}

func (handler *Handler) Upload(c echo.Context) error {
	typeFile := c.FormValue("type")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("public/uploads/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, utils.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data: echo.Map{
			"name": file.Filename,
			"url":  "localhost:8000/uploads/" + file.Filename,
			"size": file.Size,
			"type": typeFile,
		},
	})
}
