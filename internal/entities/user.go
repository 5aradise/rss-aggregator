package entities

import (
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ConvertDbUser(dbUser db.User) User {
	return User{
		dbUser.ID,
		dbUser.CreatedAt,
		dbUser.UpdatedAt,
		dbUser.Name,
	}
}

func ConvertDbUsers(dbUsers []db.User) []User {
	users := make([]User, 0, len(dbUsers))
	for _, dbUser := range dbUsers {
		users = append(users, ConvertDbUser(dbUser))
	}
	return users
}
