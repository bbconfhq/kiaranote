package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/bbconfhq/kiaranote/internal/datasource/mysql"
	"github.com/bbconfhq/kiaranote/internal/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title kiaranote
// @version 0.1
// @description kiaranote is simple notion-like service
// @termsOfService http://swagger.io/terms/

// @contact.name Ryo
// @contact.url https://github.com/bbconfhq/kiaranote
// @contact.email gwanryo@gmail.com

// @license.name MIT
// @license.url https://github.com/bbconfhq/kiaranote/blob/main/LICENSE

// @host
// @BasePath /v1
func main() {
	readerConfig := mysql.Config{
		User:   os.Getenv("READER_DB_USER"),
		Passwd: os.Getenv("READER_DB_PASS"),
		Host:   os.Getenv("READER_DB_HOST"),
		Port:   os.Getenv("READER_DB_PORT"),
		DBName: os.Getenv("READER_DB_NAME"),
	}
	reader := mysql.NewMySQL(readerConfig)
	writerConfig := mysql.Config{
		User:   os.Getenv("WRITER_DB_USER"),
		Passwd: os.Getenv("WRITER_DB_PASS"),
		Host:   os.Getenv("WRITER_DB_HOST"),
		Port:   os.Getenv("WRITER_DB_PORT"),
		DBName: os.Getenv("WRITER_DB_NAME"),
	}
	writer := mysql.NewMySQL(writerConfig)
	dao.InitRepo(reader, writer)

	e := echo.New()
	secret := os.Getenv("COOKIE_SECRET")
	e.Use(middleware.Logger())
	if os.Getenv("ENV") == "local" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
			AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowCredentials: true,
		}))
	}
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Validator = &handler.RequestValidator{
		Validator: validator.New(),
	}
	e.Use(middleware.Recover())
	handler.InitV1Handler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
