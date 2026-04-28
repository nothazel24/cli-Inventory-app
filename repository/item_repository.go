package repository

import "inventory-app/entity"

// contract
type ItemRepository interface {
	AddItem(item *entity.Item) error
	GetAllItem() ([]*entity.Item, error)
	FindItemByName(name string) (*entity.Item, error)
	Update(item *entity.Item) error
	DeleteItem(name string) error
}