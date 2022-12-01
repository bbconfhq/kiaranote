package main

import (
	"fmt"
	"github.com/bbconfhq/kiaranote/internal/dao"
	"github.com/bbconfhq/kiaranote/internal/datasource/mysql"
	"github.com/bbconfhq/kiaranote/internal/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
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
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		panic(err)
	}

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
	secret := ""
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Validator = &handler.RequestValidator{
		Validator: validator.New(),
	}
	handler.InitV1Handler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
