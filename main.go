package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Mp3         string   `xml:"guid"`
	Description string   `xml:"description"`
}

var useDir string

func main() {

	app := fiber.New()

	useDir = "./"

	if len(os.Args) > 1 {
		useDir = os.Args[1]
	}

	fmt.Println("dir", useDir)

	app.Get("/podcastapi", all)

	app.Get("/podcastapi/:id", id)

	log.Fatal(app.Listen(":7654"))

}
