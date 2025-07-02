package main

import (
	"fmt"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/application"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigServerChi{
		ServerAddress:          ":8080",
		LoaderFilePathProducts: "docs/db/products.json",
	}
	app := application.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
