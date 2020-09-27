package user

import (
	"time"

	"github.com/aslrousta/plug/lib/errors"
)

func serviceError(f, msg string, cause error) *errors.ServiceError {
	return errors.Service("user", f, msg, cause)
}

// Find retrieves a user profile. If no profile exists, it will be created.
func Find(storer Storer, id string) (*User, error) {
	var user User
	if err := storer.Find(&user, id); err != nil {
		if !errors.IsNotFound(err) {
			return nil, serviceError("Find", "failed to find user", err)
		}
		user.ID = id
		user.CreatedAt = time.Now()
		if err := storer.Save(&user); err != nil {
			return nil, serviceError("Find", "failed to save user", err)
		}
	}
	return &user, nil
}

// Update updates a user profile. If no profile exists, it will be created.
func Update(storer Storer, id string, update *UpdateRequest) (*User, error) {
	var user User
	if err := storer.Find(&user, id); err != nil {
		if !errors.IsNotFound(err) {
			return nil, serviceError("Update", "failed to find user", err)
		}
		user.ID = id
		user.CreatedAt = time.Now()
	}
	user.Nickname = update.Nickname
	user.Email = update.Email
	user.FullName = update.FullName
	user.Bio = update.Bio
	if err := storer.Save(&user); err != nil {
		return nil, serviceError("Update", "failed to save user", err)
	}
	return &user, nil
}
