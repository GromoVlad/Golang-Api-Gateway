package main

import "gin_tonic/routes"

// @securityDefinitions.apikey BearerToken
// @in header
// @name BearerToken
func main() {
	routes.Run()
}
