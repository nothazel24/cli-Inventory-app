package repository

import (
	"errors"
	"inventory-app/entity"
)

// menyimpan data dalam memori (slice)
type ItemRepositoryMemory struct {
	items []*entity.Item
}

// inisialisasi data
func NewItemRepository() ItemRepository {
	return &ItemRepositoryMemory{
		items: []*entity.Item{
			{ID: 1, Nama: "Lorem", Stok: 4, Harga: 10000},
			{ID: 2, Nama: "Indomie", Stok: 4, Harga: 3000},
		},
	}
}

func (i *ItemRepositoryMemory) AddItem(item *entity.Item) error {
	if item == nil {
		return errors.New("Item tidak boleh kosong!")
	}

	i.items = append(i.items, item)
	return nil
}

func (i *ItemRepositoryMemory) GetAllItem() ([]*entity.Item, error) {
	var result []*entity.Item

	for _, item := range i.items {
		result = append(result, item)
	}
	return result, nil
}

func (i *ItemRepositoryMemory) FindItemByName(name string) (*entity.Item, error) {
	for _, item := range i.items {
		if item.Nama == name {
			return item, nil
		}
	}
	return nil, errors.New("Item tidak ditemukan")
}

func (i *ItemRepositoryMemory) UpdateStock(name string, qty int) error {
	for _, item := range i.items {
		if item.Nama == name {
			item.Stok += qty
			return nil
		}
	}
	return errors.New("Item tidak dapat ditemukan")
}

func (i *ItemRepositoryMemory) DeleteItem(name string) error {
	for index, item := range i.items {
		if item.Nama == name {
			i.items = append(i.items[:index], i.items[index+1:]...)
			return nil
		}
	}
	return errors.New("Item tidak dapat ditemukan")
}
