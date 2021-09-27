package main

func main() {
	getEnvVariables()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(HTTP_LISTEN)
}
