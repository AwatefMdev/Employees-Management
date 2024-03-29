package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

        "github.com/AwatefMdev/graduation_project/repositories"
	"github.com/AwatefMdev/graduation_project/requests"
	"github.com/AwatefMdev/graduation_project/utils/caching"
)

type EmployeeController struct {
	DB    *sql.DB
	Cache caching.Cache
}

func NewEmployeeController(db *sql.DB, c caching.Cache) *EmployeeController {
	return &EmployeeController{
		DB:    db,
		Cache: c,
	}
}

func (jc *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	token := r.Header.Get("token")
	employeeIDStr, err := jc.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	employeeID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Convert user id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var cjr requests.CreateJobRequest
	err = decoder.Decode(&cjr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = repositories.CreateEmployee(jc.DB, cjr.Firstname, cjr.Lastname, cjr.email, cjr.adress, cjr.gender, cjr.idleaves, 
					     cjr.idtools, cjr.idattendance ,employeeID)
	if err != nil {
		log.Fatalf("Creating an employee: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (jc *EmployeeController)Employee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	EmployeeID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	employee, err := repositories.GetEmployeeByID(jc.DB, EmployeeID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
	}
	token := r.Header.Get("token")
	employeeIDStr, err := jc.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		log.Fatalf("Convert employee id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if employeeID != employee.EmployeeID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var ujr requests.UpdateJobRequest
		err = decoder.Decode(&ujr)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = repositories.UpdateJob(jc.DB, employee.ID, jc.DB, ujr.Firstname, ujr.Lastname, ujr.email, ujr.adress, ujr.gender, ujr.idleaves, 
					     ujr.idtools, ujr.idattendance )
		if err != nil {
			log.Fatalf("Updating a job: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "DELETE" {
		err = repositories.DeleteJob(jc.DB, employee.ID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

func (jc *EmployeeController) Feed(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	page := 1
	pageStr, ok := r.URL.Query()["page"]
	if ok {
		page, err = strconv.Atoi(pageStr[0])
		if err != nil {
			page = 1
		}
	}

	resultsPerPage := 10
	resultsPerPageStr, ok := r.URL.Query()["results_per_page"]
	if ok {
		resultsPerPage, err = strconv.Atoi(resultsPerPageStr[0])
		if err != nil {
			resultsPerPage = 1
		}
	}
	employees, err := repositories.GetEmployees(jc.DB, page, resultsPerPage)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}
