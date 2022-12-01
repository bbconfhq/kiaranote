package handler

import "github.com/labstack/echo/v4"

// Healthcheck   godoc
// @Summary      Health check
// @Description  Check health
// @Tags         healthcheck
// @Produce      plain
// @Success      200	{string}    string  "ok"
// @Router       /healthcheck [get]
func Healthcheck(c echo.Context) error {
    return c.String(200, "ok")
}
