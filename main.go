package main

import (
	"os"
	"strings"

	"github.com/devgony/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	const saveName string = "jobs.csv"
	defer os.Remove(scrapper.TempFileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrap(term)
	return c.Attachment(scrapper.TempFileName, saveName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
