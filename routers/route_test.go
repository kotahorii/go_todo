package routers

import (
	"bytes"
	"encoding/json"
	"go/todo3/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// TODO: Createで新しいデータを作って最後のデータと一致するか確認
func TestGetTodos(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.POST("/todos", CreateTodo)
	e.GET("/todos", GetTodos)

	param := Todo{
		Title:  "test",
		Detail: "test",
	}
	jsonParam, _ := json.Marshal(param)

	postReq, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(jsonParam))
	postReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	todos := []Todo{}
	database.DB.Find(&todos)
	jsonData, _ := json.Marshal(todos)
	expected := string(jsonData) + "\n"

	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestCreateTodo(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.POST("/todos", CreateTodo)

	param := Todo{
		Title:  "test",
		Detail: "test",
	}
	jsonParam, _ := json.Marshal(param)

	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(jsonParam))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	todo := Todo{}
	database.DB.Last(&todo)
	jsonData, _ := json.Marshal(todo)
	expected := string(jsonData) + "\n"

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestUpdateTodo(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.POST("/todos", CreateTodo)
	e.PUT("/todos/:id", UpdateTodo)

	param := Todo{
		Id:     999,
		Title:  "test",
		Detail: "test",
	}
	jsonParam, _ := json.Marshal(param)

	postReq, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(jsonParam))
	postReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	param = Todo{
		Title:  "test update",
		Detail: "test update",
	}
	jsonParam, _ = json.Marshal(param)

	req, _ := http.NewRequest(http.MethodPut, "/todos/999", bytes.NewBuffer(jsonParam))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	todo := Todo{}
	database.DB.Last(&todo)
	jsonData, _ := json.Marshal(todo)
	expected := string(jsonData) + "\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestDeleteTodo(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.DELETE("/todos/:id", DeleteTodo)

	req, _ := http.NewRequest(http.MethodDelete, "/todos/3", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
