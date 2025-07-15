package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	dummyKey := "lol"
	os.Setenv("X_API_KEY", dummyKey)

	cfg, _ := Load()

	// Ensure correct value
	assert.Equal(t, cfg.XApiKey, dummyKey)
}
