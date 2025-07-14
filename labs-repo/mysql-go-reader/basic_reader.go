package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func readFromMySQL(db *sql.DB) string {
	rows, err := db.Query("select id, name, description from items")
	if err != nil {
		log.Fatal(err)
	}

	var (
		id          int
		name        string
		description string
	)

	// TODO probably use a json library
	test := "["
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(id, name, description)
		test = test + "{ \"name\": \"" + name + "\", \"description\": \"" + description + "\"},"
	}

	test = test[:len(test)-1]

	test = test + "]"

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return test
}

// Create the database and add some default items
func createDatabase(db *sql.DB) {
	db.Query("CREATE TABLE items (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(100) NULL, description VARCHAR(1000) NULL, PRIMARY KEY (id));")
	db.Query("INSERT INTO items (name, description) VALUES (\"Garlic\", \"Roast me, roast me, say that you'll roast me\");")
	db.Query("INSERT INTO items (name, description) VALUES (\"Tomatillo\", \"Green berry\");")
}

func main() {
	fmt.Println("Welcome to the Templates Lab. First, let's check the environment variables.")

	username, usernamePresent := os.LookupEnv("MYSQL_USER")
	password, passwordPresent := os.LookupEnv("MYSQL_PASSWORD")
	// Use the environment variables supplied by
	host, hostPresent := os.LookupEnv("MYSQL_SERVICE_HOST")
	port, portPresent := os.LookupEnv("MYSQL_SERVICE_PORT")
	domain := host + ":" + port
	database, databasePresent := os.LookupEnv("MYSQL_DATABASE")

	if !usernamePresent {
		log.Fatal("Please supply a MYSQL_USER environment variable.")
	}

	if !passwordPresent {
		log.Fatal("Please supply a MYSQL_PASSWORD environment variable.")
	}

	if !hostPresent {
		log.Fatal("Please supply a MYSQL_SERVICE_HOST environment variable.")
	}

	if !portPresent {
		log.Fatal("Please supply a MYSQL_SERVICE_PORT environment variable.")
	}

	if !databasePresent {
		log.Fatal("Please supply a MYSQL_DATABASE environment variable.")
	}
	connection_string := username + ":" + password + "@tcp(" + domain + ")/" + database

	fmt.Println("connection_string = " + connection_string)

	// Open up our database connection.
	db, err := sql.Open("mysql", connection_string)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	once := true

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "false")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		if once {
			fmt.Println("Writing table and database.")
			createDatabase(db)
			once = false
		}

		fmt.Fprintf(w, readFromMySQL(db))
	})

	log.Println("About to listen on port 8080")

	http.ListenAndServe(":8080", nil)
}

