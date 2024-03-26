package main

import (
	"middleware/internal/pkg/app"
)

func main() {
	a := app.New()

	a.MustRun()
}
