package budgest

import (
	"errors"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AmbitiList(c *fiber.Ctx) error {
	db := common.GetDB()
	var ambiti []Ambito
	var count int64
	// var ambitoFilter = Ambito{IsActive: true}
	err := db.Where("budgest_ambito.is_active = true").Joins("Owner").Order("budgest_ambito.num").Find(&ambiti).Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"results": AmbitiSerializer(ambiti),
		"count":   count,
	})
}

func AmbitiDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var ambito Ambito
	id := c.Params("id")
	// var ambitoFilter = Ambito{IsActive: true}
	err := db.Where("budgest_ambito.is_active = true").Joins("Owner").First(&ambito, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(AmbitoSerializer(ambito))
}
