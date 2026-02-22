package handlers

import (
	"encoding/json"
	"go-api-app/src/database"
	"go-api-app/src/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FetchAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := database.GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

func FetchEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	employee, err := database.GetEmployeeByID(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if employee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertedID, err := database.InsertEmployee(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"status":      "success",
		"inserted_id": insertedID,
	})
}

func PatchEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alreadyExsistingEmployee, _ := database.GetEmployeeByID(idInt)
	if alreadyExsistingEmployee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	
	var emp models.Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateEmployee(idInt, emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
	})
}

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	alreadyExsistingEmployee, _ := database.GetEmployeeByID(idInt)
	if alreadyExsistingEmployee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	err = database.DeleteEmployee(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
	})
}