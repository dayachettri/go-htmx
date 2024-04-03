package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dayachettri/go-htmx/views"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Count int
}

var (
	visitCount int
)

func main() {
	e := echo.New()

	e.GET("/", homeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func homeHandler(c echo.Context) error {
	visitCount++
	currentCount := visitCount
	data := &Data{Count: currentCount}
	return render(c, http.StatusOK, views.Home(data.Count))
}

func render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response())
}
