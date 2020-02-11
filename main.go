package main

import (
	"os"
	"strings"

	"github.com/8luebottle/go-jobScrapper/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	println("Clicked")
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.Static("/assets", "assets")
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1400")) // Start a server
}
