package main

import (
	"os"
	"strings"

	"github.com/Ho-yeong/web-scrapper/scrapper"
	"github.com/labstack/echo"
)

var FileName = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrapper(c echo.Context) error {
	defer os.Remove(FileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(FileName, FileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrapper", handleScrapper)
	e.Logger.Fatal(e.Start(":1323"))
}
