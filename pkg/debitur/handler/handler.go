package handler

import (
	"mulfinance/pkg/debitur/ports"
	"mulfinance/pkg/utils"
	"net/http"
	"strconv"

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
	group.GET("/debiturs", handler.ListDebitur)
	group.GET("/debiturs/:id", handler.GetDebiturByID)
	group.GET("/debiturs/:id/limits", handler.GetDebiturLimit)
	group.GET("/debiturs/:id/transactions", handler.GetDebiturTransaction)
}

func (u *Handler) ListDebitur(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := u.usecase.ListDebitur(ctx)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}

func (u *Handler) GetDebiturByID(c echo.Context) error {
	ctx := c.Request().Context()
	userID, _ := strconv.Atoi(c.Param("id"))

	resp, err := u.usecase.GetDebiturByID(ctx, uint(userID))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}

func (u *Handler) GetDebiturLimit(c echo.Context) error {
	ctx := c.Request().Context()
	userID, _ := strconv.Atoi(c.Param("id"))

	resp, err := u.usecase.GetDebiturLimit(ctx, uint(userID))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}

func (u *Handler) GetDebiturTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	userID, _ := strconv.Atoi(c.Param("id"))

	resp, err := u.usecase.GetDebiturTransaction(ctx, uint(userID))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}
