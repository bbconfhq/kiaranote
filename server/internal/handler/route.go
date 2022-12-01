package handler

import (
	_ "github.com/bbconfhq/kiaranote/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitV1Handler(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("v1/healthcheck", Healthcheck)
	e.GET("v1/auth/login", BaseHandler(LoginRequest{}, V1Login))
	e.GET("v1/auth/logout", BaseHandler(LogoutRequest{}, V1Logout))
}
