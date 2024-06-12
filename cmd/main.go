package main

import "task-mate/internal/infrastructure/router"

func main() {
	// setup database
	// Setup and run router
	r := router.Route()
	r.Run(":8080")
}
