package enigma

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRotor(t *testing.T) {
	r := newRotor(time.Now().Unix(), math.MaxUint8)

	require.Equal(t, 'a', r.getReverse(r.getStraight('a')))

	for i := 0; i < 10; i++ {
		r.rotate()

		require.Equal(t, 'a', r.getReverse(r.getStraight('a')))
	}
}
