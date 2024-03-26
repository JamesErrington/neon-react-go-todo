package main

import (
	"net/http"

	"github.com/jameserrington/neon-arbor/app/actions"
	"github.com/jameserrington/neon-arbor/app/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "build")

	storage.InitDB()
	defer storage.CloseDB()

	e.POST("/todos", CreateTodo)

	e.Logger.Fatal(e.Start(":8080"))
}

func CreateTodo(c echo.Context) error {
	todo, err := actions.CreateTodo("Test")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}
