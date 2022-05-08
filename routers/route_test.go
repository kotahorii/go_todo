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

func TestGetTodos(t *testing.T) {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/todos", GetTodos)

	const expected = `[{"id":1,"title":"test","detail":"test","created_at":"2022-05-05T15:28:28+09:00"}]
`

	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
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
	expected := `{"id":1,"title":"test","detail":"test","created_at":"2022-05-05T15:28:28+09:00"}
`

	req, _ := http.NewRequest(http.MethodPut, "/todos/1", bytes.NewBuffer(jsonParam))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
