package main

import "kai-shop-be/internal/app"

func main() {
	// load .env automatically in db.InitPostgres
	app.Run()
}
