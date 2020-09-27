package memory

import (
	"time"

	"github.com/aslrousta/plug/lib/errors"
	"github.com/aslrousta/plug/pkg/user"
)

// UserStorage is an in-memory and unsafe storage for user profiles.
type UserStorage struct {
	Users []user.User
}

// Find finds a user by its ID.
func (s *UserStorage) Find(u *user.User, id string) error {
	index := s.indexByPred(func(u *user.User) bool {
		return u.ID == id
	})
	if index < 0 {
		return errors.NotFound("user not found", id)
	}
	*u = s.Users[index]
	return nil
}

// Save stores a user and its changes.
func (s *UserStorage) Save(u *user.User) error {
	index := s.indexByPred(func(other *user.User) bool {
		return other.ID == u.ID
	})
	if index < 0 {
		u.UpdatedAt = time.Now()
		s.Users = append(s.Users, *u)
		return nil
	}
	if s.Users[index].UpdatedAt != u.UpdatedAt {
		return errors.Integrity("user already updated", true)
	}
	s.Users[index].Nickname = u.Nickname
	s.Users[index].Email = u.Email
	s.Users[index].FullName = u.FullName
	s.Users[index].Bio = u.Bio
	s.Users[index].UpdatedAt = time.Now()
	return nil
}

func (s *UserStorage) indexByPred(pred func(*user.User) bool) int {
	for index := range s.Users {
		if pred(&s.Users[index]) {
			return index
		}
	}
	return -1
}
