package main

import "gin_tonic/routes"

// @securityDefinitions.apikey Authorization
// @in header
// @name Bearer
func main() {
	routes.Run()
}
