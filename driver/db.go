package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("DB_URL")
	fmt.Println("Connection string:", connStr)
	fmt.Println("Waiting for Database connection Start up...")
	time.Sleep(5 * time.Millisecond)

	var err error
	// ✅ Assign to the global DB, not local db
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Error opening Database: %v", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Error connecting to the Database: %v", err)
		return
	}

	fmt.Println("✅ Successfully connected to the database")
}

func GetDB() *sql.DB {
	return DB
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("❌ Error Closing The Database: %v", err) // ✅ Use lowercase %v
	}
}
