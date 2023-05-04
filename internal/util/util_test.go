package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	resp := IsValidPassword("Shizzle7*")

	assert.Equal(t, true, resp)
}

func TestGenerateAccountNumber(t *testing.T) {
	got := GenerateAccountNumber()
	want := GenerateAccountNumber()

	assert.Equal(t, want, got)
}

func TestHashPassword(t *testing.T) {
	_, err := HashPassword("shizzle")

	assert.NoError(t, err)
}

func TestInsufficientBalance(t *testing.T) {
	ok := InsufficientBalance(100, 200)

	assert.Equal(t, true, ok)
}

func TestIsValidEmail(t *testing.T) {
	ok := IsValidEmail("promis@land.com")

	assert.Equal(t, true, ok)
}
