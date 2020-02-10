package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var password string // := "secret"
	var cost int        // cost := 10
	flag.StringVar(&password, "p", "password-to-hash", "Provide the password to hash")
	flag.IntVar(&cost, "c", 10, "hashing cost of the bcrypt function")
	flag.Parse()

	if password == "password-to-hash" {
		log.Fatal("Please provide a password to hash")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
		// panic(err.Error)
	}

	fmt.Println("Your hashed password is:")
	fmt.Println(string(hashedPassword))

}
