package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusInternalServerError).SendString("Error:Sorry. Can't do it bro.")
			return nil
		},
	})

	app.Use(recover.New())
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TTTerminator by Timotej Avsec")
	})

	app.Get("/move", func(c *fiber.Ctx) error {

		board := parseBoard(c.Query("moves"))
		player := rune(c.Query("playing")[0])

		row, col := 0, 0
		if !isEmpty(board) {
			row, col = findBestMove(board, player)
		}

		sugar.Infow("New request",
			"gid", c.Query("gid"),
			"size", c.Query("size"),
			"moves", c.Query("moves"),
			"playing", c.Query("playing"),
			"move", fmt.Sprintf("%s-%d-%d", string(player), row, col),
		)

		return c.SendString(fmt.Sprintf("Move:%s-%d-%d", string(player), row, col))
	})

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
