package main

import (
	"inventory-app/cli"
	"inventory-app/repository"
	"inventory-app/service"
)

func main() {
	repo := repository.NewItemRepository()
	service := service.NewItemService(repo)

	cli.Start(service)
}