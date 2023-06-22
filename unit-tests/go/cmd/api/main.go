package main

import (
	"github.com/martinezhenry/common-exercises/unit-tests/go/internal/application"
)

func main() {
	// Initialize the application
	app := application.NewApplication()

	// Run the application
	if err := app.Run(); err != nil {
		panic(err)
	}
}
