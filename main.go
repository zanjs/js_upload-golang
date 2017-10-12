package main

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/jinzhu/configor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	// Read form fields
	// name := c.FormValue("name")
	// email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("C:/Users/Anla-E/go/src/anla.io/js_upload/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields .</p>", file.Filename))
}

// Config is ...
var Config = struct {
	APP struct {
		Name string
		Port string `default:"11323"`
	}
	Upload struct {
		Path string `default:"./"`
	}
}{}

func main() {

	configor.Load(&Config, "config.yml")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":" + Config.APP.Port))
}
