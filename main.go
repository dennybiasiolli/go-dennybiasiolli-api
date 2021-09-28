package main

import "github.com/dennybiasiolli/go-dennybiasiolli-api/common"

func main() {
	common.GetEnvVariables()

	common.ConnectDb()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(common.HTTP_LISTEN)
}
