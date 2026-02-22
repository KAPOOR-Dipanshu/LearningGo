package handlers

import (
	"encoding/json"
	"go-api-app/src/database"
	"go-api-app/src/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// FetchAllEmployees handles GET /employees and returns all employees.
func FetchAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := database.GetAllEmployees(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

// FetchEmployeeById handles GET /employees/{id} and returns a single employee.
func FetchEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	employee, err := database.GetEmployeeByID(r.Context(), idInt)
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

// CreateEmployee handles POST /employees and creates a new employee.
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertedID, err := database.InsertEmployee(r.Context(), emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"status":      "success",
		"inserted_id": insertedID,
	})
}

// PatchEmployee handles PATCH /employees/{id} and updates an employee.
func PatchEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alreadyExsistingEmployee, _ := database.GetEmployeeByID(r.Context(), idInt)
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
	err = database.UpdateEmployee(r.Context(), idInt, emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
	})
}

// RemoveEmployee handles DELETE /employees/{id} and deletes an employee.
func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVars := mux.Vars(r)
	id := pathVars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alreadyExsistingEmployee, _ := database.GetEmployeeByID(r.Context(), idInt)
	if alreadyExsistingEmployee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	err = database.DeleteEmployee(r.Context(), idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
	})
}
