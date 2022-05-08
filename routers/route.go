package routers

import (
	"go/todo3/database"
	"net/http"

	"github.com/labstack/echo"
)

type Todo struct {
	Id     int    `json:"id" param:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func GetTodos(c echo.Context) error {
	todos := []Todo{}
	database.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	database.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	todo := Todo{}
	id := c.Param("id")
	data := Todo{}
	if err := c.Bind(&data); err != nil {
		return err
	}
	database.DB.First(&todo, id).Updates(&data)
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&Todo{}, id)
	return c.NoContent(http.StatusNoContent)
}
