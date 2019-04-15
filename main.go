package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbname = "DBNAME"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
)

// main function for handle all of server request
func main() {
	initDB()
	defer db.Close()
	router := http.NewServeMux()
	http.HandleFunc("/api/index", indexHandler)
	http.HandleFunc("/api/repository", repoHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

// indexHandler does list all repos
func indexHandler(w http.ResponseWriter, r *http.Request) {}

// repoHandler does
func repoHandler(w http.ResponseWriter, r *http.Request) {}

// initDB does ...
func initDB() {
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config[dbhost], config[dbport], config[dbuser], config[dbpass], config[dbname])
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}

// dbConfig does ...
func dbConfig() map[string]string {
	conf := make(map[string]string)

	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment is required but was not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment is required but was not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment is required but was not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment is required but was not set")
	}
	pass, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment is required but was not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbname] = name
	conf[dbuser] = user
	conf[dbpass] = pass
	return conf
}
