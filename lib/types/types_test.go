package types_test

import (
	"testing"

	"github.com/aslrousta/plug/lib/types"

	. "github.com/stretchr/testify/assert"
)

func TestCellphone_IsValid(t *testing.T) {
	t.Run("Standard", func(t *testing.T) {
		phone := types.Cellphone("09121234567")
		True(t, phone.IsValid())
	})

	t.Run("ZeroLeadingInternational", func(t *testing.T) {
		phone := types.Cellphone("00989121234567")
		True(t, phone.IsValid())
	})

	t.Run("International", func(t *testing.T) {
		phone := types.Cellphone("+989121234567")
		True(t, phone.IsValid())
	})
}

func TestCellphone_Standard(t *testing.T) {
	t.Run("Standard", func(t *testing.T) {
		phone := types.Cellphone("09121234567")
		Equal(t, "09121234567", phone.Standard().String())
	})

	t.Run("ZeroLeadingInternational", func(t *testing.T) {
		phone := types.Cellphone("00989121234567")
		Equal(t, "09121234567", phone.Standard().String())
	})

	t.Run("International", func(t *testing.T) {
		phone := types.Cellphone("+989121234567")
		Equal(t, "09121234567", phone.Standard().String())
	})
}

func TestCellphone_International(t *testing.T) {
	t.Run("Standard", func(t *testing.T) {
		phone := types.Cellphone("09121234567")
		Equal(t, "+989121234567", phone.International().String())
	})

	t.Run("ZeroLeadingInternational", func(t *testing.T) {
		phone := types.Cellphone("00989121234567")
		Equal(t, "+989121234567", phone.International().String())
	})

	t.Run("International", func(t *testing.T) {
		phone := types.Cellphone("+989121234567")
		Equal(t, "+989121234567", phone.International().String())
	})
}
