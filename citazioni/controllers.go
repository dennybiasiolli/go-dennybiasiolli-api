package citazioni

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CitazioniList(c *fiber.Ctx) error {
	db := common.GetDB()
	var citazioni []Citazione
	var count int64
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).Find(&citazioni).Order("-id").Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"results": CitazioniSerializer(citazioni),
		"count":   count,
	})
}

func CitazioneDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var citazione Citazione
	id := c.Params("id")
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).First(&citazione, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(CitazioneSerializer(citazione))
}

func CitazioneCreate(c *fiber.Ctx) error {
	input := new(CreateCitazioneInput)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userTrack, _ := json.MarshalIndent(fiber.Map{
		"User-Agent":      c.Request().Header.UserAgent(),
		"x-forwarded-for": c.Get(fiber.HeaderXForwardedFor),
		// "Remote-Addr": c.Request.Header.Get("Remote-Addr"),
		"ClientIP": c.IP(),
	}, "", "  ")
	userTrackStr := string(userTrack[:])
	citazione := Citazione{
		Frase:         input.Frase,
		Autore:        input.Autore,
		UserTrackJson: &userTrackStr,
		CreatedDate:   time.Now(),
		ModifiedDate:  time.Now(),
	}
	db := common.GetDB()
	err := db.Create(&citazione).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	if common.SEND_EMAIL_AFTER_CITAZIONE_ADDED {
		go SendMailOnQuoteAdded(citazione)
	}
	return c.JSON(citazione)
}

func CitazioneRandomDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var count int64
	var citazione Citazione
	search := c.Query("search", "")
	qs := db.Model(&Citazione{}).Where(&Citazione{IsApproved: true, IsPubblica: true}).Where("frase ILIKE ? OR autore ILIKE ?", "%"+search+"%", "%"+search+"%").Count(&count)
	err := qs.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	} else if count == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "no quotes found"})
	}
	i := rand.Int63n(count)
	qs.Limit(1).Offset(int(i)).First(&citazione)
	return c.JSON(CitazioneSerializer(citazione))
}
