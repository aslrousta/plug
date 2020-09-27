package user_test

import (
	"testing"

	"github.com/aslrousta/plug/lib/rand"
	"github.com/aslrousta/plug/pkg/user"
	"github.com/aslrousta/plug/pkg/user/memory"

	. "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		id := rand.String(8)
		u, err := user.Find(&memory.UserStorage{}, id)
		if NoError(t, err) {
			Equal(t, id, u.ID)
			Empty(t, u.Nickname)
			Empty(t, u.Email)
			Empty(t, u.FullName)
			Empty(t, u.Bio)
		}
	})

	t.Run("Existing", func(t *testing.T) {
		id := rand.String(8)
		storage := &memory.UserStorage{}
		u1, _ := user.Find(storage, id)
		u2, err := user.Find(storage, id)
		if NoError(t, err) {
			Equal(t, u1.ID, u2.ID)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		id := rand.String(8)
		u, err := user.Update(&memory.UserStorage{}, id, &user.UpdateRequest{
			Nickname: "alirus",
			Email:    "aslrousta@gmail.com",
			FullName: "Ali AslRousta",
			Bio:      "Developer",
		})
		if NoError(t, err) {
			Equal(t, id, u.ID)
			Equal(t, "alirus", u.Nickname)
			Equal(t, "aslrousta@gmail.com", u.Email)
			Equal(t, "Ali AslRousta", u.FullName)
			Equal(t, "Developer", u.Bio)
		}
	})

	t.Run("Existing", func(t *testing.T) {
		id := rand.String(8)
		storage := &memory.UserStorage{}
		_, _ = user.Update(storage, id, &user.UpdateRequest{})
		u, err := user.Update(storage, id, &user.UpdateRequest{
			Nickname: "alirus",
			Email:    "aslrousta@gmail.com",
			FullName: "Ali AslRousta",
			Bio:      "Developer",
		})
		if NoError(t, err) {
			Equal(t, id, u.ID)
			Equal(t, "alirus", u.Nickname)
			Equal(t, "aslrousta@gmail.com", u.Email)
			Equal(t, "Ali AslRousta", u.FullName)
			Equal(t, "Developer", u.Bio)
		}
	})
}
