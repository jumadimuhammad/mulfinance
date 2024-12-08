package handler

import (
	"mulfinance/pkg/transaction/domain"
	"mulfinance/pkg/transaction/handler/dto"
	"mulfinance/pkg/transaction/ports"
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
	group.GET("/transactions", handler.ListTransaction)
	group.POST("/transactions", handler.CreateTransaction)
}

func (u *Handler) ListTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := u.usecase.ListTransaction(ctx)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.Response{Code: http.StatusOK, Message: "success", Data: resp})
}

func (u *Handler) CreateTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(dto.CreateTransaction)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	data := domain.Transaction{
		DebiturID:         req.DebiturID,
		DebiturLimitID:    req.DebiturLimitID,
		OTR:               req.OTR,
		AdminFee:          req.AdminFee,
		InstallmentAmount: req.InstallmentAmount,
		AmountInterest:    req.AmountInterest,
		ProductName:       req.ProductName,
		CreatedBy:         1,
		UpdatedBy:         1,
	}

	err := u.usecase.CreateTransaction(ctx, data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, utils.Response{Code: http.StatusCreated, Message: "success"})
}
