package main

import (
	"net/http"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/articoli"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/auth"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/budgest"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/citazioni"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	var db = make(map[string]string)

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	basicAuthHandler := gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", basicAuthHandler)

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	articoli.ArticoliAnonymousRegister(r.Group("/articoli"))
	citazioni.CitazioniAnonymousRegister(r.Group("/citazioni"))
	citazioni.CitazioneAnonymousRegister(r.Group("/citazione"))
	budgest.BudgestRegister(r.Group("/budgest", auth.DjangoJwtAuth))
	auth.JwtTokenRegister(r.Group("/token"))

	return r
}
