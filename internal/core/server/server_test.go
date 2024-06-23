package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {

	server := NewServer()

	assert.NotNil(t, server)
}

func TestNewServerWithPort(t *testing.T) {

	port := 5000
	server := NewServer(WithPort(port))

	assert.NotNil(t, server)
	assert.Equal(t, port, server.Port)
}
