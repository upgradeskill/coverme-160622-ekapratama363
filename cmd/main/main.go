package main

import (
	"main/internal"
	"main/internal/storage"
)

func main() {
	storage.Seeder()
	internal.SetRoutes()
}
