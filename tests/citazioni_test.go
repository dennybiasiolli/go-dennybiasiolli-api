package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/citazioni"
	"github.com/gofiber/fiber/v2"
)

func Test_citazioni(t *testing.T) {
	app, db := SetupTests(func(app *fiber.App) {
		citazioni.CitazioniAnonymousRegister(app.Group("/citazioni"))
		citazioni.CitazioneAnonymousRegister(app.Group("/citazione"))
	}, citazioni.Citazione{})

	citazione := citazioni.Citazione{
		Frase:  "Foo",
		Autore: "Bar",
	}
	resp, _ := NewJsonRequest(app, "POST", "/citazioni/", citazione)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatal("KO", resp.StatusCode)
	}

	db.First(&citazione)
	if !(citazione.Frase == "Foo" &&
		citazione.Autore == "Bar" &&
		citazione.IsApproved == false &&
		citazione.IsPubblica == false) {
		t.Fatal("Created citazione not corresponding to expected POST value")
	}

	citazione.IsApproved = true
	citazione.IsPubblica = true
	db.Updates(&citazione)

	resp, _ = NewJsonRequest(app, "GET", "/citazione/", nil)
	if resp.StatusCode == fiber.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &citazione)
		if !(citazione.Frase == "Foo" && citazione.Autore == "Bar") {
			t.Fatal()
		}
	}

}
