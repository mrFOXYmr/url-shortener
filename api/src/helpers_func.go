package src

import (
	"math/rand"
	"time"
)


func Gen_random_string() (string){
	const alf = "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lenght = 8

	// i know that random based on time isnt so random, but in this case it isnt nessosory
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, lenght)
	for i := range result{
		result[i] = alf[random.Intn(len(alf))]
	}

	return string(result)
}