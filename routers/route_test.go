package routers

import (
	"bytes"
	"encoding/json"
	"go/todo3/database"
	"go/todo3/utils"
	"net/http"
	"net/http/httptest"
	"strings"
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

	e.GET("/todos", GetTodos)

	const expected = `[{"id":1,"title":"test","detail":"test"}]
`

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

	todos := []Todo{}
	database.DB.Last(&todos)
	jsonData, _ := json.Marshal(todos)
	str := string(jsonData) + "\n"
	replacedStr := strings.Replace(str, "[", "", 1)
	expected := utils.Reverse(strings.Replace(utils.Reverse(replacedStr), "]", "", 1))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestUpdateTodo(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.PUT("/todos/:id", UpdateTodo)

	param := Todo{
		Title:  "test",
		Detail: "test",
	}
	jsonParam, _ := json.Marshal(param)
	expected := `{"id":1,"title":"test","detail":"test"}
`

	req, _ := http.NewRequest(http.MethodPut, "/todos/1", bytes.NewBuffer(jsonParam))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

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
