package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_XOR(t *testing.T) {
	cases := []struct {
		b1 uint8
		b2 uint8
		r  uint8
	}{
		{b1: 1, b2: 1, r: 0},
		{b1: 0, b2: 0, r: 0},
		{b1: 1, b2: 0, r: 1},
		{b1: 0, b2: 1, r: 1},
	}

	for _, c := range cases {
		r := XOR(c.b1, c.b2)

		assert.Equal(t, c.r, r)
	}
}

func Test_OneTimePad(t *testing.T) {
	msg := "nelson"
	key := "lsibna"

	prod := Cript(msg, key)

	assert.NotEmpty(t, prod)
	assert.NotEqual(t, msg, prod)
	assert.Equal(t, msg, Decript(key, prod))
}
