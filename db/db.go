package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB
var ErrNoUser = errors.New("no user found")

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Database connected")

	err = createTable()
	if err != nil {
		log.Fatal(err)
	}

}
func createTable() error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			age INT
		)
	`)
	return err
}
func InsertUser(name string, age int) (*User, error) {

	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	result, err := db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, name, age)
	if err != nil {
		fmt.Print(err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user := &User{
		ID:   int(lastID),
		Name: name,
		Age:  age,
	}

	return user, nil
}
func GetUserByID(userID int) (*User, error) {
	var user User

	// Use QueryRow to retrieve a single row
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Name, &user.Age)

	// Check for errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoUser
		}
		return nil, err
	}

	return &user, nil
}
func DeleteUserByID(userID int) error {

	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)

	if err != nil {
		return err
	}

	return err
}

func UpdateUser(id int, name string, age int) (*User, error) {

	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	_, err := db.Exec(`UPDATE users SET name=?, age=? WHERE id=?`, name, age, id)
	if err != nil {
		fmt.Print(err)
	}

	user := &User{
		ID:   id,
		Name: name,
		Age:  age,
	}

	return user, nil
}
