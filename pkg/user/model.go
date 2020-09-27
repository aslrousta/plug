package user

import "time"

// User is a registered user.
type User struct {
	ID        string
	Nickname  string
	Email     string
	FullName  string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storer defines a storage for users.
type Storer interface {

	// Find finds a user by its ID.
	Find(user *User, id string) error

	// Save stores a user and its changes.
	Save(user *User) error
}

// UpdateRequest is an update request for user profile.
type UpdateRequest struct {
	Nickname string
	Email    string
	FullName string
	Bio      string
}
