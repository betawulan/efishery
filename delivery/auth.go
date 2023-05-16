package delivery

import (
	"net/http"
	"strings"

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
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password, err := a.authService.Register(c.Request().Context(), user)
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

func (a authDelivery) validate(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "please provide token"})
	}

	tokens := strings.Split(token, " ")
	if len(tokens) < 2 {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "format token invalid"})
	}

	if tokens[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "no Bearer"})
	}

	user, err := a.authService.Validate(tokens[1])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)

}

func AddAuthRoute(authService service.AuthService, e *echo.Echo) {
	handler := authDelivery{
		authService: authService,
	}

	e.POST("/register", handler.register)
	e.POST("/login", handler.login)
}
