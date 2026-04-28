package service

import (
	"errors"
	"inventory-app/entity"
	"inventory-app/repository"
)

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) repository.ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) Save(item *entity.Item) error {
	if item.Nama == "" {
		return errors.New("Nama item tidak boleh kosong!")
	}
	if item.Stok < 0 {
		return errors.New("Stok tidak boleh kurang dari 0!")
	}
	if item.Harga <= 0 {
		return errors.New("Harga harus di set lebih dari 0!")
	}

	return s.repo.AddItem(item)
}

func (s *itemService) Get() ([]*entity.Item, error) {
	return s.repo.GetAllItem()
}

func (s *itemService) GetByName(name string) (*entity.Item, error) {
	if name == "" {
		return nil, errors.New("Nama tidak boleh kosong!")
	}

	return s.repo.FindItemByName(name)
}

func (s *itemService) Update(name string, qty int) error {
	if qty == 0 {
		return errors.New("Jumlah update tidak boleh 0!")
	}

	_, err := s.repo.FindItemByName(name)
	if err != nil {
		return errors.New("Item tidak ditemukan. Tidak bisa melakukan update")
	}

	return s.repo.UpdateStock(name, qty)
}

func (s *itemService) Delete(name string) error {
	if name == "" {
		return errors.New("Nama tidak boleh kosong!")
	}
	return s.repo.DeleteItem(name)
}
