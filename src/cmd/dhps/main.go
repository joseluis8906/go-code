package main

import "go.uber.org/fx"

func main() {
	app := fx.New()

	if err := app.Err(); err != nil {
		panic(err)
	}

	app.Run()
}
