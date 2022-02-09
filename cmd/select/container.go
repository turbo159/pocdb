package main

import (
	"database/sql"
	"fmt"
	"log"
)

//
func createContainerTable() (err error) {

	return err
}

// loadContainers is a utility to load a slice of container.
func loadContainers(qResults *sql.Rows) (colContaner []container, err error) {
	var item container
	for qResults.Next() {
		// for each row, scan the result into our tag composite object
		err = qResults.Scan(
			&item.SysId,
			&item.Id,
			&item.Value,
			&item.CreateUser,
			&item.CreateDate,
			&item.UpdateUser,
			&item.UpdateDate)
		if err != nil {
			log.Println("In error for loadContainers")
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		colContaner = append(colContaner, item)
	}
	return colContaner, err
}

func getCount() (count int, err error) {
	err = db.QueryRow("SELECT COUNT(id) AS count FROM container").Scan(&count)
	return count, err
}

// getContainer returns 1 container based on the id
func getContainer(id string) (item container, err error) {
	q := "SELECT * FROM container WHERE id = ?"
	results, err := db.Query(q, id)
	if err != nil {
		log.Println("getContainer...")
		return item, fmt.Errorf("no records found")
	}

	containerCol, err := loadContainers(results)
	if len(containerCol) < 1 {
		return item, err
	} else {
		return containerCol[0], err
	}
}

// getAllContainers returns a container slice of all containers.
func getAllContainer() (containerCol []container, err error) {
	results, err := db.Query("SELECT * FROM container")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	containerCol, err = loadContainers(results)

	return containerCol, err
}

func deleteContainer(id string) (code int64, err error) {

	result, _ := db.Exec("DELETE FROM container WHERE id = ?", id)
	code, err = result.RowsAffected()
	return code, err
}

func insertContainer(item container) (code int64, err error) {
	stmt, err := db.Prepare("insert into container(id, value, createuser, updateuser) values (?, ?, ?, ?)")
	ErrorCheck(err)

	res, err := stmt.Exec(item.Id, item.Value, item.CreateUser, item.UpdateUser)
	ErrorCheck(err)

	code, err = res.LastInsertId()
	ErrorCheck(err)

	return code, err
}
