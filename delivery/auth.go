package delivery

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/betawulan/efishery/error_message"
	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/service"
)

type authDelivery struct {
	authService service.AuthService
}

type credential struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type successLogin struct {
	Token string `json:"token"`
}

func (a authDelivery) register(c echo.Context) error {
	var register model.User

	err := c.Bind(&register)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password, err := a.authService.Register(c.Request().Context(), register)
	if err != nil {
		switch _err := err.(type) {
		case error_message.Duplicate:
			return c.JSON(http.StatusConflict, _err)
		}

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, password)
}

func (a authDelivery) login(c echo.Context) error {
	cred := credential{}

	err := c.Bind(&cred)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := a.authService.Login(c.Request().Context(), cred.Phone, cred.Password)
	if err != nil {
		switch _err := err.(type) {
		case error_message.Unauthorized:
			return c.JSON(http.StatusUnauthorized, _err)

		case error_message.NotFound:
			return c.JSON(http.StatusNotFound, _err)
		}

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, successLogin{Token: token})
}

func AddAuthRoute(authService service.AuthService, e *echo.Echo) {
	handler := authDelivery{
		authService: authService,
	}

	e.POST("/register", handler.register)
	e.POST("/login", handler.login)
}
