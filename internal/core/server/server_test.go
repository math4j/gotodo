package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {

	server := NewServer()

	assert.NotNil(t, server)
}

func TestNewServerWithPort(t *testing.T) {

	port := 5000
	addr := fmt.Sprintf(":%d", port)
	server := NewServer(WithPort(port))

	assert.NotNil(t, server)
	assert.Equal(t, addr, server.Addr)
}
