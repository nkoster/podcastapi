package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func all(c *fiber.Ctx) error {

	files, err := ioutil.ReadDir(useDir)

	if err != nil {
		log.Fatal(err)
	}

	jsonString := "{\"data\":["
	for _, file := range files {
		name := file.Name()
		ext := name[len(name)-4:]
		name = name[:len(name)-4]
		if ext == ".xml" {
			jsonString = jsonString + "{\"name\":\"" +
				name + "\"},"
		}
	}
	jsonString = strings.TrimRight(jsonString, ",")
	jsonString += "]}"

	c.Set("Content-type", "application/json; charset=utf-8")

	return c.SendString(jsonString)

}
