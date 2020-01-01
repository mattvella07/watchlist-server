package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// DB is the reference to the database
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=db port=%d user=%s password=%s dbname=%s sslmode=disable", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME")))
	if err != nil {
		log.Fatalf("Unable to connect to DB: %s\n", err)
	}

	retries := 15
	for retries >= 0 {
		err = DB.Ping()
		if err != nil {
			if retries == 0 {
				log.Fatalf("Unable to communicate with DB: %s\n", err)
			}

			log.Print("Error communicating with DB, trying again...")
			time.Sleep(1 * time.Second)

			retries--
		} else {
			retries = -1
		}
	}
}
