package handle

import (
	"math/rand"
	"time"
)

type rnd struct {
	RandomNumber int `json:"rnd"`
}

func RandomNumGen() rnd {
	rand.Seed(time.Now().UnixNano())
	rndNum := rand.Intn(26)

	var randomNumber rnd
	randomNumber.RandomNumber = rndNum
	return randomNumber

}
