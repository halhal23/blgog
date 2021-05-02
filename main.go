package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()
var db *sqlx.DB

const templatePath = "src/template/"

func main() {
	db = connectDB()
	e.GET("/", articleIndex)
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/:id/edit", articleEdit)

	e.Logger.Fatal(e.Start(":8080"))
}

////////////////
///設定関数
////////////////

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("connect db successfully")
	return db
}

////////////////
///ハンドラ関数
////////////////

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"message": "Hello Hello desune",
		"time":    time.Now(),
	}
	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"message": "hello new page",
		"time":    time.Now(),
	}
	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"message": "show page",
		"time":    time.Now(),
		"id":      id,
	}
	return render(c, "article/show.html", data)
}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"message": "edit page",
		"time":    time.Now(),
		"id":      id,
	}
	return render(c, "article/edit.html", data)
}

////////////////
///プライベート関数
////////////////

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
