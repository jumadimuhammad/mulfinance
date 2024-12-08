package handler

import (
	"mulfinance/pkg/auth/handler/dto"
	"mulfinance/pkg/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Routes(group *echo.Group) {
	group.POST("/login", handler.Login)
}

func (u *Handler) Login(c echo.Context) error {
	req := new(dto.Login)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	if !utils.ComparePassword(os.Getenv("LOGIN_PASSWORD"), req.Password) || req.Username != os.Getenv("LOGIN_USERNAME") {
		return echo.ErrUnauthorized
	}

	claims := &Claims{
		os.Getenv("LOGIN_NAME"),
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
