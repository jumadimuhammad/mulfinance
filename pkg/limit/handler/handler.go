package handler

import (
	"mulfinance/pkg/limit/ports"
	"mulfinance/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase ports.IUsecase
}

func NewHandler(usecase ports.IUsecase) *Handler {
	return &Handler{
		usecase,
	}
}

func (handler *Handler) Routes(group *echo.Group) {
	group.GET("/limits", handler.ListLimit)
}

func (u *Handler) ListLimit(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := u.usecase.ListLimit(ctx)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}
