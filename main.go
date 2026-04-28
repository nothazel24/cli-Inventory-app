package main

import (
	"fmt"
	"inventory-app/cli"
	"inventory-app/entity"
	"inventory-app/repository"
	"inventory-app/service"
)

func main() {
	repo := repository.NewItemRepository()
	service := service.NewItemService(repo)

	cli.ShowMenu()

	err := service.Save(&entity.Item{ID: 1, Nama: "Aqua", Stok: 4, Harga: 4000})
	if err != nil {
		fmt.Println(err.Error())
	}

	items, _ := service.Get()
	for _, item := range items {
		fmt.Println(item.Nama, "-", item.Stok)
	}
}