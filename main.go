package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

type User struct {
	IDUser      int    `json:"idUser"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type Post struct {
	IDPost      int    `json:"idPost"`
	ContentText string `json:"content_text"`
	CreatedAt   string `json:"created_at"`
	UserID      int    `json:"userID"`
}

type Comment struct {
	IDComment   int    `json:"idComment"`
	IDPost      int    `json:"idPost"`
	IDUser      int    `json:"idUser"`
	ContentText string `json:"content_text"`
	CreatedAt   string `json:"created_at"`
}

func GetAllPosts(c echo.Context) error {
	// Get all posts from the database
	query := `SELECT idPost, content_text, created_at, userID FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to query posts"})
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.IDPost, &post.ContentText, &post.CreatedAt, &post.UserID); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan post data"})
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error iterating over rows"})
	}

	// Return posts as JSON response
	return c.JSON(http.StatusOK, posts)
}

func GetPostByUserID(c echo.Context) error {
	userID := c.QueryParam("id")

	// Get all posts from the database
	query := `SELECT idPost, content_text, created_at, userID FROM posts WHERE userID = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to query posts"})
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.IDPost, &post.ContentText, &post.CreatedAt, &post.UserID); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan post data"})
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error iterating over rows"})
	}

	// Return posts as JSON response
	return c.JSON(http.StatusOK, posts)
}

func GetAllCommentsToPost(c echo.Context) error {
	postID := c.QueryParam("idPost")

	// Get all comments from the database
	query := `SELECT idComment, idPost, idUser, content_text, created_at FROM comments WHERE idPost = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to query comments"})
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.IDComment, &comment.IDPost, &comment.IDUser, &comment.ContentText, &comment.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan comment data"})
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error iterating over rows"})
	}

	// Return comments as JSON response
	return c.JSON(http.StatusOK, comments)
}

func GetUserByID(c echo.Context) error {
	userID := c.QueryParam("id")

	// Get user from the database
	query := `SELECT idUser, username, displayName, email FROM users WHERE idUser = ?`
	row := db.QueryRow(query, userID)

	var user User
	if err := row.Scan(&user.IDUser, &user.Username, &user.DisplayName, &user.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan user data"})
	}

	// Return user as JSON response
	return c.JSON(http.StatusOK, user)
}


func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	db = database

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

	// Start the server
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"}, // Adjust as necessary
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/posts", GetAllPosts)
	e.GET("/comments", GetAllCommentsToPost)
	e.GET("/users", GetUserByID)
	e.GET("/posts/user", GetPostByUserID)
	e.Logger.Fatal(e.Start(":5050"))

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
