package domain

import "errors"

var ErrItemNotFound = errors.New("Item not found")

type ItemRepositoryStub struct {
	items []Item
}

func (s ItemRepositoryStub) FindAll() ([]Item, error) {
	return s.items, nil
}

func (s ItemRepositoryStub) InsertItem(item Item) error {
	// Implement code to insert the item into the stub repository
	s.items = append(s.items, item)
	return nil
}

func (s ItemRepositoryStub) FindItemByID(itemID string) (Item, error) {
	for _, c := range s.items {
		if c.ItemID == itemID { // Use ItemID instead of Id
			return c, nil
		}
	}
	return Item{}, ErrItemNotFound
}

func (s ItemRepositoryStub) UpdateItemByID(itemID string, updatedItem Item) error {
	for i, c := range s.items {
		if c.ItemID == itemID { // Use ItemID instead of Id
			s.items[i] = updatedItem
			return nil
		}
	}
	return ErrItemNotFound
}

func NewItemRepositoryStub() ItemRepositoryStub {
	items := []Item{
		{"1001", "Yamada", "Kanagawa", 1110000, "1970-01-01", "1"}, // Use integer values for Quantity
		{"1002", "Akiyama", "Akita", 2220000, "1960-01-01", "2"},   // Use integer values for Quantity
	}
	return ItemRepositoryStub{items: items}
}
