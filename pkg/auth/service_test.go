package auth_test

import (
	"testing"

	"github.com/aslrousta/plug/pkg/auth"
	"github.com/aslrousta/plug/pkg/auth/memory"

	. "github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	t.Run("InvalidPhone", func(t *testing.T) {
		_, err := auth.SignIn(&memory.AccountStorage{}, "1234567")
		EqualError(t, err, "phone number is invalid")
	})

	t.Run("NewUser", func(t *testing.T) {
		acc, err := auth.SignIn(&memory.AccountStorage{}, "09121234567")
		if NoError(t, err) {
			NotEmpty(t, acc.ID)
			NotEmpty(t, acc.RefreshToken)
			Equal(t, "+989121234567", acc.Phone.String())
		}
	})

	t.Run("ExistingUser", func(t *testing.T) {
		storage := &memory.AccountStorage{}
		_, _ = auth.SignIn(storage, "09121234567")
		_, err := auth.SignIn(storage, "09121234567")
		NoError(t, err)
	})
}
