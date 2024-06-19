package main

import (
	"log"

	"github.com/aliazizii/url-shotener/internal/database"
	"github.com/aliazizii/url-shotener/internal/handler"
	"github.com/aliazizii/url-shotener/internal/model"
	"github.com/aliazizii/url-shotener/internal/store"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Url{})
	if err != nil {
		log.Fatal(err)
	}

	s := store.NewSQL(db)
	urlHandler := handler.Url{
		Store: s,
	}

	app.POST("/insert", urlHandler.Insert)
	app.POST("/find-by-hash", urlHandler.FindByHash)

	if err := app.Start(":1234"); err != nil {
		log.Fatal("cannot start the http server")
	}
}
