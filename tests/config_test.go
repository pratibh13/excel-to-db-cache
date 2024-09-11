package tests

import (
	"assignment/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("InitDB caused panic: %v", r)
		}
	}()
	config.InitDB()
	assert.NotNil(t, config.DB, "DB connection should be initialized")
}

func TestInitRedis(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("InitRedis caused panic: %v", r)
		}
	}()
	config.InitRedis()
	assert.NotNil(t, config.RDB, "Redis client should be initialized")
}
