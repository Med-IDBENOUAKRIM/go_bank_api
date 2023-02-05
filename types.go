package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAccount(first_name, last_name string) *Account {
	return &Account{
		ID:        rand.Int(),
		FirstName: first_name,
		LastName:  last_name,
		Number:    int64(rand.Int()),
	}
}
