/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 22:57
 * Copyright (c) 2022
 */

package helper

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

const (
	AlphaNumericLowUp = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRandomPassword() (string, error) {
	bytePass := []byte(RandString(10))
	password, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return string(password), err
}

func RandString(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = AlphaNumericLowUp[rand.Intn(len(AlphaNumericLowUp))]
	}

	return string(b)
}
