package main

import (
	"math/rand"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	common.GetEnvVariables()

	common.ConnectDb()
	// common.GetDB().AutoMigrate(&citazioni.Citazione{})

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(common.HTTP_LISTEN)
}
