package handler

import (
	"net/http"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/labstack/echo/v4"
)

const templatePath = "src/template/"

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(templatePath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
