package memory

import (
	"time"

	"github.com/aslrousta/plug/lib/errors"
	"github.com/aslrousta/plug/lib/types"
	"github.com/aslrousta/plug/pkg/auth"
)

// AccountStorage is an in-memory and unsafe storage for user accounts.
type AccountStorage struct {
	Accounts []auth.Account
}

// Find finds an account by its ID.
func (s *AccountStorage) Find(account *auth.Account, id string) error {
	index := s.indexByPred(func(account *auth.Account) bool {
		return account.ID == id
	})
	if index < 0 {
		return errors.NotFound("account not found", id)
	}
	*account = s.Accounts[index]
	return nil
}

// FindByPhone finds an account by its phone number.
func (s *AccountStorage) FindByPhone(account *auth.Account, phone types.Cellphone) error {
	index := s.indexByPred(func(account *auth.Account) bool {
		return account.Phone == phone
	})
	if index < 0 {
		return errors.NotFound("account not found", string(phone))
	}
	*account = s.Accounts[index]
	return nil
}

// FindByRefreshToken finds an account by its refresh token.
func (s *AccountStorage) FindByRefreshToken(account *auth.Account, refreshToken string) error {
	index := s.indexByPred(func(account *auth.Account) bool {
		return account.RefreshToken == refreshToken
	})
	if index < 0 {
		return errors.NotFound("account not found", refreshToken)
	}
	*account = s.Accounts[index]
	return nil
}

// Save stores an account and its changes.
func (s *AccountStorage) Save(account *auth.Account) error {
	index := s.indexByPred(func(other *auth.Account) bool {
		return other.ID == account.ID
	})
	if index < 0 {
		account.UpdatedAt = time.Now()
		s.Accounts = append(s.Accounts, *account)
		return nil
	}
	if s.Accounts[index].UpdatedAt != account.UpdatedAt {
		return errors.Integrity("account already updated", true)
	}
	s.Accounts[index].RefreshToken = account.RefreshToken
	s.Accounts[index].UpdatedAt = time.Now()
	return nil
}

func (s *AccountStorage) indexByPred(pred func(*auth.Account) bool) int {
	for index := range s.Accounts {
		if pred(&s.Accounts[index]) {
			return index
		}
	}
	return -1
}
