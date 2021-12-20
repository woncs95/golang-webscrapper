package main

import (
	"fmt"
	"os"
	"scrapper"
	"strings"

	"github.com/labstack/echo/v4"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	defer func() {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}()
	return c.Attachment(fileName, fileName)
}

// func getenv(key, fallback string) string {
//     value := os.Getenv(key)
//     if len(value) == 0 {
//         return fallback
//     }
//     return value
// }

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	
	// serverport := getenv("SERVER_PORT", "1323") 
	// portString := fmt.Sprintf(":%s", serverport)

	e.Logger.Fatal(e.Start(":1323"))
	// if err := e.Start(portString); err != http.ErrServerClosed {
	// 	e.Logger.Fatal(err)
	// }
}
