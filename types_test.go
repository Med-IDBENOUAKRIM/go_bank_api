package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account, err := NewAccount("med", "gear", "gearsnakeman")
	assert.Nil(t, err)

	fmt.Printf("%+v", account)
}
