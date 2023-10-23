package domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type itemRepositoryDb struct {
	db *gorm.DB
}

func (d itemRepositoryDb) FindAll() ([]Item, error) {
	var items []Item
	if err := d.db.Find(&items).Error; err != nil {
		log.Println("Error while querying items table: " + err.Error())
		return nil, err
	}
	return items, nil
}

func (d itemRepositoryDb) InsertItem(item Item) error {
	if err := d.db.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func (d itemRepositoryDb) FindItemByID(itemID string) (Item, error) {
	var item Item
	if err := d.db.Where("item_id = ?", itemID).First(&item).Error; err != nil {
		log.Println("Error while querying item by ID: " + err.Error())
		return Item{}, err
	}
	return item, nil
}

func (d itemRepositoryDb) UpdateItemByID(itemID string, updatedItem Item) error {
	if err := d.db.Model(&Item{}).Where("item_id = ?", itemID).Updates(updatedItem).Error; err != nil {
		log.Println("Error while updating item by ID: " + err.Error())
		return err
	}
	return nil
}

func NewItemRepositoryDb() itemRepositoryDb {
	db, err := gorm.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Item{})

	return itemRepositoryDb{db}
}

func (d itemRepositoryDb) DeleteItemByID(itemID string) error {
	if err := d.db.Where("item_id = ?", itemID).Delete(&Item{}).Error; err != nil {
		log.Println("Error while deleting item by ID: " + err.Error())
		return err
	}
	return nil
}

