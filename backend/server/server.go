package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"unhaus/model"
	"unhaus/utils"
)

func redirect(c *fiber.Ctx) error {
	unhausUrl := c.Params("redirect")
	unhaus, err := model.FindByUnhausUrl(unhausUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not find goly in DB " + err.Error(),
		})
	}

	unhaus.Clicked +=1
	err = model.UpdateUnhaus(unhaus)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}
	return c.Redirect(unhaus.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllUnhaus(c *fiber.Ctx) error {
	unhaus, err := model.GetAllUnhaus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all unhaus links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(unhaus)

}

func getUnhaus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}
	unhaus, err := model.GetUnhaus(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could no retreive unahus from db " + err.Error(), 
		})
	}
	return c.Status(fiber.StatusOK).JSON(unhaus)
}

func createUnhaus(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var unhaus model.Unhaus
	err := c.BodyParser(&unhaus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}
	if unhaus.Random {
		unhaus.Unhaus = utils.RandomURL(10)
	}
	err = model.CreateUnhaus(unhaus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create unhaus in db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(unhaus)
}

func updateUnhaus(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var unhaus model.Unhaus
	err := c.BodyParser(&unhaus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}
	err = model.UpdateUnhaus(unhaus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update unhaus link in DB " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(unhaus)
}

func deleteUnhaus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}
	err = model.DeleteUnhaus(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "unhaus deleted",
	})
}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Get("/r/:redirect", redirect)
	router.Get("/unhaus", getAllUnhaus)
	router.Get("/unhaus/:id", getUnhaus)
	router.Post("/unhaus", createUnhaus)
	router.Patch("/unhaus", updateUnhaus)
	router.Delete("/unhaus/:id", deleteUnhaus)
	router.Listen(":3001")
}