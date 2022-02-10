package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var (
	db    *sql.DB
	dbErr error
)

func init() {
	usr := os.Getenv("DB_USR")
	pwd := os.Getenv("DB_PWD")
	host := os.Getenv("DB_HOST")

	if usr != "" && pwd != "" && host != "" {
		conString := fmt.Sprint(usr, ":", pwd, "@tcp(", host, ")/gm_testing")
		log.Println("Connection: ", conString)
		db, dbErr = sql.Open("mysql", conString)
		ErrorCheck(dbErr)
	} else {
		panic("Missing environment variables: DB_USR, DB_PWD, DB_HOST")
	}
}

func main() {
	defer db.Close()

	var version string

	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	ErrorCheck(err)

	fmt.Println(version)

	var newItem container
	rand.Seed(time.Now().UnixNano())
	newItem.Id = uuid.New().String()
	newItem.Value = "What is going on? - " + strconv.Itoa(rand.Int())
	newItem.CreateUser = "gmizuno"

	log.Println("insertContainer")
	result, err := insertContainer(newItem)
	ErrorCheck(err)
	log.Println(result)

	results, err := getAllContainer()
	ErrorCheck(err)

	// and then print out the tag's Name attribute
	for i, item := range results {
		log.Println(i, item.SysId, item.Id, item.Value, item.CreateUser, item.CreateDate)
	}

	log.Println("deleteContainer")
	result, err = deleteContainer("6fc41c5f-8154-4eb8-b0e7-43bc7e40409e")
	ErrorCheck(err)
	log.Println("Delete code: ", result)

	//log.Println("getContainer")
	// item, err := getContainer("89067ecc-6f7a-4032-9837-9f0d1972367b")
	// ErrorCheck(err)
	// log.Println("getContainer: ", item.Id, item.Value, item.CreateUser, item.CreateDate)
	log.Println("count")
	count, err := getCount()
	ErrorCheck(err)
	log.Println("Total records:", count)
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal("There were crappy errors. \n", err.Error())
		panic(err.Error())
	}
}
