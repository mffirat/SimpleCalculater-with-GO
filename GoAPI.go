package main

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func API() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		aStr := c.Query("a")
		bStr := c.Query("b")
		op := c.Query("op")

		a, err := strconv.ParseFloat(aStr, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid Number: " + aStr})
		}
		b, err := strconv.ParseFloat(bStr, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid Number: " + bStr})
		}
		var result float64
		if op == "add" {
			result = a + b
		} else if op == "sub" {
			result = a - b
		} else if op == "mul" {
			result = a * b
		} else if op == "div" {
			if b == 0 {
				return c.Status(400).JSON(fiber.Map{"error": "A number dont divide with 0"})
			}
			result = a / b
		} else if op == "pow" {
			result = math.Pow(a, b)
		} else if op == "root" {
			if b == 0 {
				return c.Status(400).JSON(fiber.Map{"error": "A numbers  does not root  with 0"})
			}
			result = math.Pow(a, 1/b)
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid operation"})
		}
		return c.Status(400).JSON(fiber.Map{"result: ": result})

	})

	app.Listen(":8080")

}
