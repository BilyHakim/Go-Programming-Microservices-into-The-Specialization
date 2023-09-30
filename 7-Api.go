package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employess = []Employee{
	{ID: 1, Name: "Bily", Age: 23, Division: "IT"},
	{ID: 2, Name: "Hakim", Age: 33, Division: "Finance"},
	{ID: 3, Name: "Erlangga", Age: 19, Division: "Marketing"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listeing on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//
	//if r.Method == "GET" {
	//	json.NewEncoder(w).Encode(employess)
	//	return
	//}

	tpl, err := template.ParseFiles("template.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, employess)
	return

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employess) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employess = append(employess, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
