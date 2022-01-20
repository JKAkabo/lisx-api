package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lisxAPI/models"
	"lisxAPI/repos"
	"log"
)

func CreateAnalyzer(c *fiber.Ctx) error {
	var analyzerCreate models.AnalyzerCreate
	if err := c.BodyParser(&analyzerCreate); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	id, err := repos.InsertAnalyzer(
		analyzerCreate.Name,
		analyzerCreate.Protocol,
		analyzerCreate.IP,
		analyzerCreate.Port,
		analyzerCreate.ServerMode,
		analyzerCreate.SerialPort,
		analyzerCreate.BaudRate,
		analyzerCreate.Parity,
		analyzerCreate.DataBits,
		analyzerCreate.StopBits,
		analyzerCreate.StartDelimiter,
		analyzerCreate.EndDelimiter,
		analyzerCreate.UserID,
	)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	analyzer, err := repos.SelectAnalyzerById(id)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(analyzer)
}

func GetAnalyzerById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	analyzer, err := repos.SelectAnalyzerById(id)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(analyzer)
}

func GetAnalyzers(c *fiber.Ctx) error {
	analyzers, err := repos.SelectAnalyzers()
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(analyzers)
}
