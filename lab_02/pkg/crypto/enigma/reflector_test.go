package enigma

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReflector(t *testing.T) {
	r := newReflector(time.Now().Unix(), 0, math.MaxUint8)

	require.Equal(t, byte('a'), r.Reflect(r.Reflect('a')))
}
