package main

import (
    "time"
    "fmt"
)

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    string
}

func (s *Server) CreateUser(
	username string,
	email string,
	passwordHash string,
) error {
	_, err := s.db.Exec(`
        INSERT INTO users (
            username,
            email,
            password_hash,
            created_at
        )
        VALUES (?, ?, ?, ?)
    `,
		username,
		email,
		passwordHash,
		time.Now().UTC().Format(time.RFC3339),
	)

	return err
}
func (s *Server) GetUserByEmail(email string) (*User, error) {
    var user User
    err:= s.db.QueryRow(`
        SELECT id, username, email, password_hash, created_at
        FROM users
        WHERE email = ?
    `, email).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.PasswordHash,
        &user.CreatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &user, nil
}
func testUserDatabase(server *Server) {
    // err := server.CreateUser(
    //     "Stanly",
    //     "stanlyvarghese16@gmail.com",
    //     "Plot no. 33 Gopal Vihar, Barra-6, Kanpur, Uttar Pradesh - 208027",
    // )

    // if err != nil {
    //     fmt.Println("CreateUser error:", err)
    //     return
    // }

    user, err := server.GetUserByEmail("stanlyvarghese14@gmail.com")

    if err != nil {
        fmt.Println("GetUserByEmail error:", err)
        return
    }
    fmt.Println("User ID:", user.ID)
    fmt.Println("Username:", user.Username)
    fmt.Println("Email:", user.Email)
}
