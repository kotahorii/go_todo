package main

import (
	"github.com/labstack/echo"

	"go/todo3/database"
	"go/todo3/routers"
)

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/todos", routers.GetTodos)
	e.POST("/todos", routers.CreateTodo)
	e.PUT("/todos/:id", routers.UpdateTodo)
	e.DELETE("/todos/:id", routers.DeleteTodo)

	e.Logger.Fatal(e.Start(":3000"))
}
