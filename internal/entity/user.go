package entity

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	ID        int64
	PubId     uuid.UUID
	Name      string
	CreatedAt time.Time
}

func NewUser(name string) *User {
	return &User{Name: strings.ToUpper(name)}
}
