package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(first_name, last_name string) *Account {
	return &Account{
		ID:        rand.Int(),
		FirstName: first_name,
		LastName:  last_name,
		Number:    int64(rand.Int()),
	}
}
