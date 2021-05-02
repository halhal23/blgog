package handler

import (
	"net/http"
	"strconv"
	"time"

	"techpit-blog/repository"

	"github.com/labstack/echo/v4"
)

func ArticleIndex(c echo.Context) error {
	articles, err := repository.ArticleList()
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}
	data := map[string]interface{}{
		"message":  "Hello Hello desune",
		"time":     time.Now(),
		"articles": articles,
	}
	return render(c, "article/index.html", data)
}

func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"message": "hello new page",
		"time":    time.Now(),
	}
	return render(c, "article/new.html", data)
}

func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"message": "show page",
		"time":    time.Now(),
		"id":      id,
	}
	return render(c, "article/show.html", data)
}

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"message": "edit page",
		"time":    time.Now(),
		"id":      id,
	}
	return render(c, "article/edit.html", data)
}
