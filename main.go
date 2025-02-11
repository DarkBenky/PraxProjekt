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
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const (
	testing = true
	testingLogin = "test"
	testingPassword = "test"
)

type User struct {
	IDUser      int    `json:"idUser"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
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

func GetUsers() []User {
	// Get all users from the database
	query := `SELECT idUser, username, displayName, email FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.IDUser, &user.Username, &user.DisplayName, &user.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func GetPosts() []Post {
	// Get all posts from the database
	query := `SELECT idPost, content_text, created_at, userID FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.IDPost, &post.ContentText, &post.CreatedAt, &post.UserID); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return posts
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

func GetAllUsers(c echo.Context) error {
	// Get all users from the database
	query := `SELECT idUser, username, displayName, email FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to query users"})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.IDUser, &user.Username, &user.DisplayName, &user.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan user data"})
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error iterating over rows"})
	}

	// Return users as JSON response
	return c.JSON(http.StatusOK, users)
}

func AddPost(c echo.Context) error {
    type PostRequest struct {
        UserID      string `json:"userID"`
        ContentText string `json:"contentText"`
    }

    post := new(PostRequest)
    if err := c.Bind(post); err != nil {
		fmt.Println(err)
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid request data",
        })
    }

    if post.UserID == "" || post.ContentText == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "UserID and content text are required",
        })
    }

    query := `INSERT INTO posts (userID, content_text, created_at) VALUES (?, ?, ?)`
    result, err := db.Exec(query, post.UserID, post.ContentText, time.Now().Format(time.RFC3339))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to insert post: " + err.Error(),
        })
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil || rowsAffected == 0 {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to confirm post insertion",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "Post added successfully",
    })
}

func AddComment(c echo.Context) error {
	type CommentRequest struct {
		PostID      string `json:"postID"`
		UserID      string `json:"userID"`
		ContentText string `json:"contentText"`
	}

	comment := new(CommentRequest)
	if err := c.Bind(comment); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request data",
		})
	}

	// Validate input
	if comment.PostID == "" || comment.UserID == "" || comment.ContentText == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "PostID, userID, and content text are required",
		})
	}

	// Insert comment into the database
	query := `INSERT INTO comments (idPost, idUser, content_text, created_at) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, comment.PostID, comment.UserID, comment.ContentText, time.Now().Format(time.RFC3339))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to insert comment: " + err.Error(),
		})
	}

	// Check if the insert was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to confirm comment insertion",
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Comment added successfully",
	})
}


// In main.go, update the EditPost function:
func EditPost(c echo.Context) error {
    // Create struct for request body
    type EditRequest struct {
        PostID      string `json:"postID"`
        ContentText string `json:"contentText"`
    }

    // Parse request body
    var req EditRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid request format",
        })
    }

    // Validate input
    if req.PostID == "" || req.ContentText == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Post ID and content text are required",
        })
    }

    // Update post in database
    query := `UPDATE posts SET content_text = ? WHERE idPost = ?`
    result, err := db.Exec(query, req.ContentText, req.PostID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to update post: " + err.Error(),
        })
    }

    // Verify update
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to confirm update",
        })
    }
    if rowsAffected == 0 {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "Post not found",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "Post updated successfully",
    })
}

// In main.go, update the DeletePost function:
func DeletePost(c echo.Context) error {
    // Create struct for request body
    type DeleteRequest struct {
        PostID string `json:"postID"`
    }

    // Parse request body
    var req DeleteRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid request format",
        })
    }

    // Validate postID
    if req.PostID == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Post ID is required",
        })
    }

    // Delete post from database
    query := `DELETE FROM posts WHERE idPost = ?`
    result, err := db.Exec(query, req.PostID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to delete post: " + err.Error(),
        })
    }

    // Verify deletion
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to confirm deletion",
        })
    }
    if rowsAffected == 0 {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "Post not found",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "Post deleted successfully",
    })
}

func GetPostById(c echo.Context) error {
	postID := c.QueryParam("id")

	// Get post from the database
	query := `SELECT idPost, content_text, created_at, userID FROM posts WHERE idPost = ?`
	row := db.QueryRow(query, postID)

	var post Post
	if err := row.Scan(&post.IDPost, &post.ContentText, &post.CreatedAt, &post.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan post data"})
	}

	// Return post as JSON response
	return c.JSON(http.StatusOK, post)
}



func main() {

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
	// n := 10 // Number of random entries to generate

	// fmt.Printf("Generating %d random users...\n", n)
	// insertRandomUsers(n)

	// n = n / 2
	// fmt.Printf("Generating %d random posts...\n", n)
	// insertRandomPosts(n)

	// n = n / 2
	// fmt.Printf("Generating %d random comments...\n", n)
	// insertRandomComments(n)

	// InsertTestUser()

	// Start the server
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:8080", "http://127.0.0.1:8080"},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowCredentials: true,
    }))

	e.GET("/posts", GetAllPosts)
	e.GET("/comments", GetAllCommentsToPost)
	e.GET("/user", GetUserByID)
	e.GET("/users", GetAllUsers)
	e.GET("/posts/user", GetPostByUserID)
	e.GET("/post", GetPostById)
	e.POST("/addPost", AddPost)
	e.POST("/addComment", AddComment)
	e.DELETE("/deletePost", DeletePost)
	e.PUT("/editPost", EditPost)
	e.POST("/login", Login)
	e.PUT("/userEdit", UpdateUser)
	e.Logger.Fatal(e.Start(":5050"))

}

func InsertTestUser() {
	_, err := db.Exec(`INSERT INTO users (username, displayName, email, password) VALUES (?, ?, ?, ?)`,
		testingLogin, testingLogin,testingLogin + "@gmail.com", testingPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted test user.\n")

	// insert random post and comment for test user
	_, err = db.Exec(`INSERT INTO posts (content_text, created_at, userID) VALUES (?, ?, ?)`,
		faker.Sentence(), time.Now().Format(time.RFC3339), 1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO comments (idPost, idUser, content_text, created_at) VALUES (?, ?, ?, ?)`,
		1, 1, faker.Sentence(), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted random post and comment for test user.\n")
}

type UpdateUserRequest struct {
    ID          int    `json:"id"`
    Username    string `json:"username"`
    DisplayName string `json:"displayName"`
    Email       string `json:"email"`
}

func UpdateUser(c echo.Context) error {
    var req UpdateUserRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid request payload",
        })
    }


    // Update the user in the database
    stmt, err := db.Prepare("UPDATE users SET username = ?, displayName = ?, email = ? WHERE idUser = ?")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Database error",
        })
    }
    defer stmt.Close()

    res, err := stmt.Exec(req.Username, req.DisplayName, req.Email, req.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to update user",
        })
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Database error",
        })
    }

    if rowsAffected == 0 {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "User not found",
        })
    }

    // Retrieve the updated user
    var updatedUser User
    err = db.QueryRow("SELECT idUser, username, displayName, email FROM users WHERE idUser = ?", req.ID).
        Scan(&updatedUser.IDUser, &updatedUser.Username, &updatedUser.DisplayName, &updatedUser.Email)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to retrieve updated user",
        })
    }

    return c.JSON(http.StatusOK, updatedUser)
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format",
		})
	}

	// Validate input
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Username and password are required",
		})
	}

	// Check if user exists
	query := `SELECT idUser, username, displayName, email FROM users WHERE (username = ? OR email = ?) AND password = ?`
	row := db.QueryRow(query, req.Username, req.Username, req.Password)

	var user User
	if err := row.Scan(&user.IDUser, &user.Username, &user.DisplayName, &user.Email); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// Create Users table
func createUsersTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"idUser" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"username" TEXT UNIQUE,
		"displayName" TEXT UNIQUE,
		"email" TEXT UNIQUE,
		"password" TEXT
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

// insertRandomUsers generates and inserts synthetic user data into the database for testing purposes.
// It creates 'n' users with randomly generated usernames, display names, and emails using the faker library.
//
// Parameters:
//   - n: The number of random users to generate and insert
//
// The function will:
//   - Generate unique random usernames, display names and emails for each user
//   - Insert the generated data into the 'users' table
//   - Print confirmation message after successful insertion
//
// Example usage:
//
//	insertRandomUsers(100) // Generates 100 random users
//
// Note: This function will fail fast with log.Fatal if any database errors occur
func insertRandomUsers(n int) {
	for i := 0; i < n; i++ {
		username := faker.Username()
		displayName := faker.Username()
		email := faker.Email()
		password := faker.Password()
		_, err := db.Exec(`INSERT INTO users (username, displayName, email, password) VALUES (?, ?, ?, ?)`,
			username, displayName, email, password)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Inserted %d random users.\n", n)
}

// insertRandomPosts generates and inserts synthetic post data for existing users in the database.
// For each user, it creates a random number of posts (1 to n) with randomly generated content.
//
// Parameters:
//   - n: The maximum number of posts to generate per user
//
// The function will:
//   - Retrieve all existing users from the database
//   - For each user, generate between 1 and n random posts
//   - Set the creation timestamp to the current time
//   - Insert the generated posts into the 'posts' table
//   - Print confirmation message after successful insertion
//
// Example usage:
//
//	insertRandomPosts(5) // Generates up to 5 posts per user
//
// Note:
//   - Requires existing users in the database
//   - Will fail fast with log.Fatal if any database errors occur
func insertRandomPosts(n int) {
	Users := GetUsers()

	for _, user := range Users {
		numberOfPosts := rand.Intn(n) + 1
		for i := 0; i < numberOfPosts; i++ {
			content := faker.Sentence()
			createdAt := time.Now().Format(time.RFC3339)
			_, err := db.Exec(`INSERT INTO posts (content_text, created_at, userID) VALUES (?, ?, ?)`,
				content, createdAt, user.IDUser)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Printf("Inserted random posts for %d users.\n", len(Users))
}

// insertRandomComments generates and inserts synthetic comment data for existing posts in the database.
// For each post, it creates a random number of comments (1 to n) with randomly generated content
// and assigns them to random users.
//
// Parameters:
//   - n: The maximum number of comments to generate per post
//
// The function will:
//   - Retrieve all existing posts and users from the database
//   - For each post, generate between 1 and n random comments
//   - Randomly assign each comment to an existing user
//   - Set the creation timestamp to the current time
//   - Insert the generated comments into the 'comments' table
//   - Print confirmation message after successful insertion
//
// Example usage:
//
//	insertRandomComments(10) // Generates up to 10 comments per post
//
// Note:
//   - Requires existing posts and users in the database
//   - Will fail fast with log.Fatal if any database errors occur
//   - Uses random user selection, so comment distribution may not be uniform
func insertRandomComments(n int) {
	Posts := GetPosts()
	Users := GetUsers()

	for _, post := range Posts {
		numberOfComments := rand.Intn(n) + 1

		for i := 0; i < numberOfComments; i++ {
			content := faker.Sentence()
			createdAt := time.Now().Format(time.RFC3339)
			_, err := db.Exec(`INSERT INTO comments (idPost, idUser, content_text, created_at) VALUES (?, ?, ?, ?)`,
				post.IDPost, Users[rand.Intn(len(Users))].IDUser, content, createdAt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Printf("Inserted random comments for %d posts.\n", len(Posts))
}
