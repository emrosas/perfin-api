package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/emrosas/perfin-api/pkg/handlers"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type User struct {
	Id            int
	Username      string
	Password      string
	Session_token string
	CSRF_token    string
}

func initDatabase() *sql.DB {
	databaseURL := os.Getenv("TURSO_DATABASE_URL")
	databaseToken := os.Getenv("TURSO_AUTH_TOKEN")
	url := databaseURL + "?authToken=" + databaseToken

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	fmt.Println("db opened")
	return db
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %s\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Session_token, &user.CSRF_token); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		users = append(users, user)
		fmt.Println(user.Id, user.Username)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := initDatabase()
	defer db.Close()
	queryUsers(db)
}
