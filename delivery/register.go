package delivery

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/packages/error_message"
	"github.com/betawulan/efishery/service"
)

type registerDelivery struct {
	registerService service.RegisterService
}

func (r registerDelivery) register(c echo.Context) error {
	var register model.Register

	err := c.Bind(&register)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password, err := r.registerService.Register(c.Request().Context(), register)
	if err != nil {
		switch _err := err.(type) {
		case error_message.Duplicate:
			return c.JSON(http.StatusConflict, _err)
		}

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, password)
}

func AddRegisterRoute(registerService service.RegisterService, e *echo.Echo) {
	handler := registerDelivery{
		registerService: registerService,
	}

	e.POST("/register", handler.register)
}
