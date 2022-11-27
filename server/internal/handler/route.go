package handler

import "github.com/labstack/echo/v4"

func InitV1Handler(e *echo.Echo) {
	e.GET("v1/healthcheck", Healthcheck)
	e.GET("v1/auth/login", BaseHandler(LoginRequest{}, V1Login))
	e.GET("v1/auth/logout", BaseHandler(LogoutRequest{}, V1Logout))
}
