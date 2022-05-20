package pi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PI(t *testing.T) {
	cases := []struct {
		i int
		r float64
	}{
		{i: 1000000, r: 3.1415916535897743},
	}

	for _, c := range cases {
		pi := CalculatePi(c.i)

		assert.Equal(t, c.r, pi)
	}
}
