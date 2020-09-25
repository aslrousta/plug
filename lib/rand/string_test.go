package rand_test

import (
	"testing"

	"github.com/aslrousta/plug/lib/rand"

	. "github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	sample := rand.String(8)
	Len(t, sample, 8)
}
