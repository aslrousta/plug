package auth

import (
	"time"

	"github.com/aslrousta/plug/lib/errors"
	"github.com/aslrousta/plug/lib/rand"
	"github.com/aslrousta/plug/lib/types"
)

const (
	accountIDLength    = 16
	refreshTokenLength = 48
)

// SignIn logs a user into the system by their phone number. If the user does
// not exist, it will be created.
func SignIn(storer AccountStorer, phone string) (*Account, error) {
	cellphone := types.Cellphone(phone).International()
	if !cellphone.IsValid() {
		return nil, errors.Validation("phone number is invalid", "phone")
	}
	var account Account
	if err := storer.FindByPhone(&account, cellphone); err != nil {
		if !errors.IsNotFound(err) {
			return nil, err
		}
		account.ID = rand.String(accountIDLength)
		account.Phone = cellphone
		account.RefreshToken = rand.String(refreshTokenLength)
		account.RegisteredAt = time.Now()
		if err := storer.Save(&account); err != nil {
			return nil, err
		}
	}
	return &account, nil
}
