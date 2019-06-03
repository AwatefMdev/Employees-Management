package repositories

import (
	"database/sql"
	"github.com/AwatefMdev/graduation_project/models"

func CreateEmployee(db *sql.DB, firstname, lastname,email,adress,gender string, idleaves,idtools,idattendance int) (int, error) {
	const query = `
		insert into employees (
			firstname,
			lastname,
			email,
			gender,
			idleaves,
			idtools,
			idattendance
			
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		) returning id
	`
	var id int
	err := db.QueryRow(query, firstname, lastname,email,adress,gender, idleaves,idtools,idattendance).Scan(&id)
	return id, err
}

func UpdateJob(db *sql.DB, idemployee int, firstname, lastname,email,adress,gender string) error {
	const query = `
		update employees set
			firstname=$1
			lastname=$2
			email=$3
			gender=$4
			idleaves=$5
			idtools=$6
			idattendance=$7
		where id = $8
	`
	_, err := db.Exec(query, firstname, lastname,email,adress,gender, idemployee)
	return err
}

func DeleteEmployee(db *sql.DB, id int) error {
	const query = `delete from jobs where id = $1`
	_, err := db.Exec(query, id)
	return err
}

func GetEmployeeByID(db *sql.DB, id int) (*models.Job, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			jobs
		where
			id = $1
	`
	var job models.Job
	err := db.QueryRow(query, id).Scan(&job.ID, &job.Title, &job.Description, &job.UserID)
	return &job, err
}

func GetEmployee(db *sql.DB, page, resultsPerPage int) ([]*models.Job, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			jobs
		limit $1 offset $2
	`
	jobs := make([]*models.Job, 0)
	offset := (page - 1) * resultsPerPage

	rows, err := db.Query(query, resultsPerPage, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var job models.Job
		err = rows.Scan(&job.ID, &job.Title, &job.Description, &job.UserID)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, err
}
