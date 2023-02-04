package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internals/core/application/handlers"
	"testing"
)

var (
	handler = URLHandlerTest()
)

func TestUrlHandlerGetAll(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/urls", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result := handler.Index(c)

	assert.NoError(t, result)
	assert.Equal(t, 200, rec.Code)
}

func URLHandlerTest() handlers.UrlHandlerInterface {
	mock := URLRepositoryMock()
	return handlers.NewUrlHandler(mock)
}
