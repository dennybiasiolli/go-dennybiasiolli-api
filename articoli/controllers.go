package articoli

import (
	"errors"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ArticoliList(c *fiber.Ctx) error {
	db := common.GetDB()
	var articoli []Articolo
	var count int64
	err := db.Where("data_pubblicazione IS NOT NULL").Find(&articoli).Order("-id").Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"results": ArticoliSerializer(articoli),
		"count":   count,
	})
}

func ArticoliDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var articolo Articolo
	id := c.Params("id")
	err := db.Where("data_pubblicazione IS NOT NULL").First(&articolo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(ArticoloSerializer(articolo))
}
