package repository

import "inventory-app/entity"

// contract
type ItemRepository interface {
	AddItem(item *entity.Item) error
	GetAllItem() ([]*entity.Item, error)
	FindItemByName(name string) (*entity.Item, error)
	UpdateStock(name string, qty int) error
	DeleteItem(name string) error
}

type ItemService interface {
	Save(item *entity.Item) error
	Get() ([]*entity.Item, error)
	GetByName(name string) (*entity.Item, error)
	Update(name string, qty int) error
	Delete(name string) error
}
