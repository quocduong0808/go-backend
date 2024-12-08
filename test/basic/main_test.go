package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOneBasic(t *testing.T) {
	out := AddOne(3, 4)
	expected := 7
	if out != 7 {
		t.Errorf("expected output %d, actual %d", expected, out)
	}
}

func TestAddOneUsingTestify(t *testing.T) {
	out := AddOne(3, 4)
	expected := 7
	assert.Equal(t, expected, out)
}
