package main

import "time"

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