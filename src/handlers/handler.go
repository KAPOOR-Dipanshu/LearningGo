package handlers

import (
	"encoding/json"
	"go-api-app/src/database"
	"go-api-app/src/models"
	"net/http"
)

var items []models.Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if items == nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Current list of items is empty",
			"status":  "success",
		})
		return
	}
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func GetDataFromMongoDB(w http.ResponseWriter, r *http.Request) {
	dbCollection := database.GetCollection()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":     "success",
		"collection": dbCollection.Name(),
	})
}
