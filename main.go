package main

import (
  "flag"
  "fmt"
  "log"
)

func seedAccount(store Storage, first_name, last_name, pw string) *Account {
  acc, err := NewAccount(first_name, last_name, pw)
  if err != nil {
    log.Fatal(err)
  }

  if err := store.CreateAccount(acc); err != nil {
    log.Fatal(err)
  }

  fmt.Println("New account => ", acc.Number)

  return acc
}

func seedAccounts(s Storage) {
  seedAccount(s, "Marcos", "MWM", "max7777")
}

func main() {
  seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
