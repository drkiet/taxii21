package main

import (
	"github.com/drkiet/accounts"
	"log"
	"testing"
)

func TestVerifyUser(t *testing.T) {
	log.Println("test Verify user")
	accounts.Config([]string{"localhost:2379", "localhost:2380"}, 5)
	var user = "test_user"
	var password = "test_password"

	if !CreateUser(user, password) {
		t.Error("test verify user create credential fails")
	}

	log.Println("verifying user ...", user, password)
	if !VerifyUser(user, password) {
		t.Error("test verify user fails")
	}
	log.Println("test Verify user ends")
}
