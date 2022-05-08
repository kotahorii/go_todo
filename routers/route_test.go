package routers

import (
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

	expected := `[{"id":1,"title":"test","detail":"test","created_at":"2022-05-05T15:28:28+09:00"}]
`

	e.GET("/todos", GetTodos)

	req, _ := http.NewRequest(http.MethodGet, "localhost:3000/todos", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
