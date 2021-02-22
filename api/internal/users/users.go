package users

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"

	database "github.com/advenjourney/api/internal/pkg/db/postgres"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create(ctx context.Context) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(ctx, "INSERT INTO Users(Username,Password) VALUES($1,$2)", user.Username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPasswordHash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// GetUserIDByUsername check if a user exists in database by given username
func GetUserIDByUsername(username string) (int, error) {
	ctx := context.Background()
	row := database.DB.QueryRow(ctx, "SELECT ID from Users WHERE Username = $1", username)

	var ID int
	if err := row.Scan(&ID); err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}

		return 0, err
	}

	return ID, nil
}

// Authenticate authenticates a user
func (user *User) Authenticate() bool {
	ctx := context.Background()
	row := database.DB.QueryRow(ctx, "SELECT Password from Users WHERE Username = $1", user.Username)
	var hashedPassword string
	if err := row.Scan(&hashedPassword); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("unexpected query error in auth: %v", err)
		}

		return false
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}
