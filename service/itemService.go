package service

import (
	"github.com/fadedreams/go_hexagonal_rest/domain"
)

type ItemService interface {
	GetAllItems() ([]domain.Item, error)
	CreateItem(domain.Item) error
	GetItemByID(itemID string) (domain.Item, error)
	UpdateItemByID(itemID string, item domain.Item) error
	DeleteItemByID(itemID string) error
}

type DefaultItemService struct {
	repo domain.ItemRepository
}

func (s DefaultItemService) GetAllItems() ([]domain.Item, error) {
	return s.repo.FindAll()
}

func (s DefaultItemService) CreateItem(item domain.Item) error {
	// Implement the logic to create a new item
	// Here, we assume that your repository has an InsertItem method
	err := s.repo.InsertItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultItemService) GetItemByID(itemID string) (domain.Item, error) {
	// Implement the logic to retrieve a item by ID from the repository
	item, err := s.repo.FindItemByID(itemID)
	if err != nil {
		return domain.Item{}, err
	}
	return item, nil
}

func (s DefaultItemService) UpdateItemByID(itemID string, updatedItem domain.Item) error {
	err := s.repo.UpdateItemByID(itemID, updatedItem)
	if err != nil {
		return err
	}
	return nil
}

func NewItemService(repository domain.ItemRepository) DefaultItemService {
	return DefaultItemService{repo: repository}
}

func (s DefaultItemService) DeleteItemByID(itemID string) error {
	err := s.repo.DeleteItemByID(itemID)
	if err != nil {
		return err
	}
	return nil
}
