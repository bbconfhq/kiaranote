package handler

import (
	_ "github.com/bbconfhq/kiaranote/docs"
	"github.com/bbconfhq/kiaranote/internal/constant"
	"github.com/bbconfhq/kiaranote/internal/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitV1Handler(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("v1/healthcheck", Healthcheck)

	e.POST("v1/auth/login", BaseHandler(LoginRequest{}, V1Login))
	e.GET("v1/auth/logout", BaseHandler(LogoutRequest{}, V1Logout))
	e.POST("v1/auth/register", BaseHandler(RegisterRequest{}, V1Register))

	e.GET("v1/user", BaseHandler(GetUsersRequest{}, V1GetUsers), middleware.AuthMiddleware(constant.RoleAdmin))
	e.POST("v1/user", BaseHandler(PostUserRequest{}, V1PostUser), middleware.AuthMiddleware(constant.RoleAdmin))
	e.GET("v1/user/:user_id", BaseHandler(GetUserRequest{}, V1GetUser), middleware.AuthMiddleware(constant.RoleUser))
	e.PUT("v1/user/:user_id", BaseHandler(PutUserRequest{}, V1PutUser), middleware.AuthMiddleware(constant.RoleUser))
	e.DELETE("v1/user/:user_id", BaseHandler(DeleteUserRequest{}, V1DeleteUser), middleware.AuthMiddleware(constant.RoleUser))

	e.GET("v1/note", BaseHandler(GetNotesRequest{}, V1GetNotes), middleware.AuthMiddleware(constant.RoleUser))
	e.POST("v1/note", BaseHandler(PostNoteRequest{}, V1PostNote), middleware.AuthMiddleware(constant.RoleAdmin))
	e.GET("v1/note/:note_id", BaseHandler(GetNoteRequest{}, V1GetNote), middleware.AuthMiddleware(constant.RoleUser))
	e.PUT("v1/note/:note_id", BaseHandler(PutNoteRequest{}, V1PutNote), middleware.AuthMiddleware(constant.RoleAdmin))
	e.DELETE("v1/note/:note_id", BaseHandler(DeleteNoteRequest{}, V1DeleteNote), middleware.AuthMiddleware(constant.RoleAdmin))
}
