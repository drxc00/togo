package utils

import (
	"math/rand"
	"time"
)

var defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateID() string {
	// Generate an ID using a random number generator
	// Uses a default alphabet of characters
	id := make([]rune, 16)
	// Fill the ID with random characters from the default alphabet
	for i := range id {
		id[i] = defaultAlphabet[rand.Intn(len(defaultAlphabet))]
	}
	return "togo_" + string(id)
}

func GetCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}
