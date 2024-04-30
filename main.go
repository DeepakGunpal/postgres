// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello from server")
// 	})

// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	})

// 	// Channel to receive server start confirmation
// 	started := make(chan bool, 3)

// 	// For server1
// 	go func() {
// 		err := http.ListenAndServe(":8080", nil)
// 		if err != nil {
// 			fmt.Println("Error starting server1:", err)
// 		} else {
// 			fmt.Println("Server1 started")
// 			started <- true
// 		}
// 	}()

// 	// For server2
// 	go func() {
// 		err := http.ListenAndServe(":8081", nil)
// 		if err != nil {
// 			fmt.Println("Error starting server2:", err)
// 		} else {
// 			fmt.Println("Server2 started")
// 			started <- true
// 		}
// 	}()

// 	// For server3
// 	go func() {
// 		err := http.ListenAndServe(":8082", nil)
// 		if err != nil {
// 			fmt.Println("Error starting server3:", err)
// 		} else {
// 			fmt.Println("Server3 started")
// 			started <- true
// 		}
// 	}()

// 	// Wait for server start confirmations
// 	for i := 0; i < 3; i++ {
// 		if <-started {
// 			fmt.Printf("Server listening on :%d\n", 8080+i)
// 		}
// 	}

// 	// Block the main goroutine indefinitely to prevent the program from exiting
// 	select {}
// }

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database credentials
	dbUser := "horocosmo"
	dbPass := "horocosmo"
	dbHost := "3.6.88.119"
	dbName := "postgres"

	// Construct the data source name
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", dbUser, dbPass, dbHost, dbName)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Ping the database to verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database!")
}
