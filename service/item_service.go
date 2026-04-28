package service

import (
	"errors"
	"inventory-app/entity"
	"inventory-app/repository"
)

type ItemService interface {
	AddItem(name string, harga, stok int) error
	GetAllItem() ([]*entity.Item, error)
	UpdateStock(name string, qty int) error
	DeleteItem(name string) error
	FindItem(name string) (*entity.Item, error)
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) AddItem(name string, harga, stok int) error {
	if name == "" {
		return errors.New("Nama item tidak boleh kosong!")
	}
	if stok < 0 {
		return errors.New("Stok tidak boleh kurang dari 0!")
	}
	if harga <= 0 {
		return errors.New("Harga harus di set lebih dari 0!")
	}

	_, err := s.repo.FindItemByName(name)
	if err == nil {
		return errors.New("item sudah ada")
	}

	item := &entity.Item{
		Nama:  name,
		Stok:  stok,
		Harga: harga,
	}

	return s.repo.AddItem(item)
}

func (s *itemService) GetAllItem() ([]*entity.Item, error) {
	return s.repo.GetAllItem()
}

func (s *itemService) UpdateStock(name string, qty int) error {
	item, err := s.repo.FindItemByName(name)
	if err != nil {
		return err
	}

	if item.Stok+qty < 0 {
		return errors.New("Jumlah stok tidak boleh negatif!")
	}

	item.Stok += qty
	return s.repo.Update(item)
}

func (s *itemService) FindItem(name string) (*entity.Item, error) {
	return s.repo.FindItemByName(name)
}

func (s *itemService) DeleteItem(name string) error {
	return s.repo.DeleteItem(name)
}
