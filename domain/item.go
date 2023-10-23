package domain

type Item struct {
	ItemID      string `gorm:"column:item_id;primary_key"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Quantity    int    `gorm:"column:quantity"`
	Color       string `gorm:"column:color"`
	Status      string `gorm:"column:status"`
}

type ItemRepository interface {
	FindAll() ([]Item, error)
	InsertItem(Item) error
	FindItemByID(itemID string) (Item, error)
	UpdateItemByID(string, Item) error
	DeleteItemByID(itemID string) error
}
