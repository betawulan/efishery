package delivery

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/betawulan/efishery/packages/error_message"
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