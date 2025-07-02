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
		ServerAddress:       ":8080",
		LoaderFilePathBuyer: "docs/db/buyers.json",
		LoaderFilePathProducts: "docs/db/products.json",
		LoaderFilePathSeller:   "docs/db/sellers.json",
		LoaderFilePathWarehouse: "docs/db/warehouse.json",
	}
	app := application.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
