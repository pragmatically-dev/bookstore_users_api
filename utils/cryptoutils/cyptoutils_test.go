package cryptoutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordhash(t *testing.T) {
	password := "santiago"
	hash, err := HashPassword(password)
	assert.Nil(t, err, "Fail")

	assert.True(t, CheckPasswordHash("santiago", hash))

	



}
