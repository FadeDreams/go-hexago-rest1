package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/fadedreams/go_hexagonal_rest/domain"
	"github.com/fadedreams/go_hexagonal_rest/service"
	"github.com/gorilla/mux"
	"log"
)

type ItemHandlers struct {
	service service.ItemService
}

func (ch *ItemHandlers) getAllItems(w http.ResponseWriter, r *http.Request) {
	items, _ := ch.service.GetAllItems()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(items)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func (ch *ItemHandlers) createItem(w http.ResponseWriter, r *http.Request) {
	var item domain.Item

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := ch.service.CreateItem(item)
	if err != nil {
		log.Printf("Failed to create item: %v", err) // Log the error
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (ch *ItemHandlers) getItemById(w http.ResponseWriter, r *http.Request) {
	// Get the item ID from the request, assuming it's in the URL parameters
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Call the service to retrieve the item by ID
	item, err := ch.service.GetItemByID(itemID)
	if err != nil {
		// Handle the error, for example, by sending an error response
		http.Error(w, "Failed to get item by ID", http.StatusInternalServerError)
		return
	}

	// Respond with the item details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (ch *ItemHandlers) updateItemById(w http.ResponseWriter, r *http.Request) {
	// Get the item ID from the URL parameters
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Decode the updated item data from the request body
	var updatedItem domain.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service to update the item by ID
	err := ch.service.UpdateItemByID(itemID, updatedItem)
	if err != nil {
		// Handle the error, for example, by sending an error response
		http.Error(w, "Failed to update item by ID", http.StatusInternalServerError)
		return
	}

	// Respond with a success status code and a confirmation message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Item updated successfully")
}

func (ch *ItemHandlers) deleteItemById(w http.ResponseWriter, r *http.Request) {
	// Get the item ID from the URL parameters
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Call the service to delete the item by ID
	err := ch.service.DeleteItemByID(itemID)
	if err != nil {
		// Handle the error, for example, by sending an error response
		http.Error(w, "Failed to delete item by ID", http.StatusInternalServerError)
		return
	}

	// Respond with a success status code and a confirmation message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Item deleted successfully")
}
