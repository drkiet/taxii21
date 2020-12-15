package main

import (
	"github.com/drkiet/accounts"
	log "log"
)

func VerifyUser(user, password string) bool {
	log.Println("verifying user ...", user, password)
	return accounts.VerifyCredential(user, password)
}

func CreateUser(user, password string) bool {
	log.Println("create user ...", user, password)
	return accounts.CreateCredential(user, password)
}