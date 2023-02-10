package main

import (
	"flag"
	"log"
)

func seedAccount(store Storage, first_name, last_name, password string) *Account {
	account, err := NewAccount(first_name, last_name, password)
	if err != nil {
		log.Fatalln(err)
	}

	if err := store.CreateAccount(account); err != nil {
		log.Fatalln("Seeding accounts got error : ", err)
	}

	return account
}

func main() {
	seed := flag.Bool("seed", false, "seed th database")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		log.Println("seeding the database")
		seedAccount(store, "Mohamed", "Idben", "pass123")
	}

	server := NewServer(":3550", store)
	server.RunServer()
}
