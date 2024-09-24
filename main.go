package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	idUser      int    `json:"idUser"`
	username    string `json:"username"`
	displayName string `json:"displayName"`
	email       string `json:"email"`
}

type Post struct {
	idPost      int    `json:"idPost"`
	contentText string `json:"content_text"`
	createdAt   string `json:"created_at"`
	userID      int    `json:"userID"`
}

type Comment struct {
	idComment   int    `json:"idComment"`
	idPost      int    `json:"idPost"`
	idUser      int    `json:"idUser"`
	contentText string `json:"content_text"`
	createdAt   string `json:"created_at"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Create tables
	createUsersTable(database)
	createPostsTable(database)
	createCommentsTable(database)

	// Generate random users, posts, and comments
	n := 50 // Number of random entries to generate

	insertRandomUsers(database, n)
	insertRandomPosts(database, n)
	insertRandomComments(database, n)

	fmt.Printf("Inserted %d random users, posts, and comments.\n", n)
}

// Create Users table
func createUsersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"idUser" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT,
		"displayName" TEXT,
		"email" TEXT
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Users table created")
}

// Create Posts table
func createPostsTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS posts (
		"idPost" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"content_text" TEXT,
		"created_at" TEXT,
		"userID" INTEGER,
		FOREIGN KEY(userID) REFERENCES users(idUser)
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Posts table created")
}

// Create Comments table
func createCommentsTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS comments (
		"idComment" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"idPost" INTEGER,
		"idUser" INTEGER,
		"content_text" TEXT,
		"created_at" TEXT,
		FOREIGN KEY(idPost) REFERENCES posts(idPost),
		FOREIGN KEY(idUser) REFERENCES users(idUser)
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Comments table created")
}

// Insert n random users
func insertRandomUsers(db *sql.DB, n int) {
	for i := 0; i < n; i++ {
		username := faker.Username()
		displayName := faker.Username()
		email := faker.Email()
		_, err := db.Exec(`INSERT INTO users (username, displayName, email) VALUES (?, ?, ?)`,
			username, displayName, email)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Inserted %d random users.\n", n)
}

// Insert n random posts
func insertRandomPosts(db *sql.DB, n int) {
	for i := 0; i < n; i++ {
		content := faker.Sentence()
		createdAt := time.Now().Format(time.RFC3339)
		userID := rand.Intn(n) + 1 // Random user ID between 1 and n

		_, err := db.Exec(`INSERT INTO posts (content_text, created_at, userID) VALUES (?, ?, ?)`,
			content, createdAt, userID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Inserted %d random posts.\n", n)
}

// Insert n random comments
func insertRandomComments(db *sql.DB, n int) {
	for i := 0; i < n; i++ {
		content := faker.Sentence()
		createdAt := time.Now().Format(time.RFC3339)
		postID := rand.Intn(n) + 1 // Random post ID between 1 and n
		userID := rand.Intn(n) + 1 // Random user ID between 1 and n

		_, err := db.Exec(`INSERT INTO comments (idPost, idUser, content_text, created_at) VALUES (?, ?, ?, ?)`,
			postID, userID, content, createdAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Inserted %d random comments.\n", n)
}
