package delivery

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/betawulan/efishery/error_message"
	"github.com/betawulan/efishery/service"
)

type fishDelivery struct {
	fishService service.FishService
	authService service.AuthService
}

//	@Summary		fetch
//	@Description	fetch resources
//	@Tags			fetch
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Success		200				{array}		[]model.Fish
//	@Failure		401				{object}	error_message.Unauthorized
//	@Router			/app [get]
func (f fishDelivery) fetch(c echo.Context) error {
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

	fishes, err := f.fishService.GetDataStorages(tokens[1])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, fishes)
}

func AddFishRoute(fish service.FishService, e *echo.Echo) {
	handler := fishDelivery{fishService: fish}

	e.GET("/app", handler.fetch)
}
