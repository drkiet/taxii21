package main

import (
	"fmt"
	"github.com/drkiet/accounts"
	log "log"
)

func VerifyUser(user, password string) bool {
	log.Println("verifying user ...", user, password)
	accounts.Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var hash, err = accounts.Get(user)
	if err != nil {
		log.Println(fmt.Printf("User %s not found", user))
		return false
	}
	if len(hash) == 0 {
		log.Println(fmt.Printf("User %s may not exist!", user))
		return false
	}

	if !accounts.ComparePasswords(hash, []byte(password)) {
		log.Println("User/password is invalid")
		return false
	}
	return true
}