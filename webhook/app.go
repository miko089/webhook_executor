package webhook

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os/exec"
	"test2/configParser"
)

func New(filename, authKey string) *fiber.App {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		payload := struct {
			AuthKey string `json:"auth_key"`
			ID      string `json:"id"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			c.Status(400)
			return c.SendString("wrong request")
		}
		if payload.AuthKey != authKey {
			c.Status(401)
			return c.SendString("wrong auth key")
		}
		command, err := configParser.GetCommand(filename, payload.ID)
		if err != nil {
			c.Status(404)
			return c.SendString("wrong id")
		}
		cmd := exec.Command("bash", "-c", command)
		err = cmd.Run()
		if err != nil {
			c.Status(400)
			return c.SendString(fmt.Sprintf("error while executing command: %+v", err))
		}
		return c.SendStatus(200)
	})
	return app
}
