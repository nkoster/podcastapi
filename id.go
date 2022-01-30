package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func id(c *fiber.Ctx) error {

	id := c.Params("id")

	data, err := os.ReadFile(useDir + id + ".xml")

	if err != nil {
		return c.SendString("not found")
	}

	var item Item
	xml.Unmarshal(data, &item)

	fmt.Println(item.Title)

	jsonString := "{\"data\":{"
	jsonString += "\"title\":\"" + item.Title + "\","
	jsonString += "\"link\":\"" + item.Mp3 + "\","
	jsonString += "\"description\":\"" + item.Description + "\""
	jsonString += "}}"

	c.Set("Content-type", "application/json; charset=utf-8")

	return c.SendString(jsonString)

}