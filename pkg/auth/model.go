package auth

import (
	"time"

	"github.com/aslrousta/plug/lib/types"
)

// Account represents a registered user.
type Account struct {
	ID           string
	Phone        types.Cellphone
	RefreshToken string
	RegisteredAt time.Time
	UpdatedAt    time.Time
}

// AccountStorer defines a storage for users.
type AccountStorer interface {

	// Find finds an account by its ID.
	Find(account *Account, id string) error

	// FindByPhone finds an account by its phone number.
	FindByPhone(account *Account, phone types.Cellphone) error

	// FindByRefreshToken finds an account by its refresh token.
	FindByRefreshToken(account *Account, refreshToken string) error

	// Save stores an account and its changes.
	Save(account *Account) error
}
