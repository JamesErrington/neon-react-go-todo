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

	e.GET("/todos", GetAllTodos)
	e.POST("/todos", CreateTodo)
	e.PUT("/todo/:id", UpdateTodo)

	e.Logger.Fatal(e.Start(":8080"))
}

type CreateBody struct {
	Task string `json:"task"`
}

func CreateTodo(c echo.Context) error {
	data := CreateBody{}
	c.Bind(&data)

	todo, err := actions.CreateTodo(data.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func GetAllTodos(c echo.Context) error {
	todos, err := actions.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todos)
}

type UpdateBody struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func UpdateTodo(c echo.Context) error {
	data := UpdateBody{}
	c.Bind(&data)

	err := actions.UpdateTodo(data.Id, data.Task, data.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
