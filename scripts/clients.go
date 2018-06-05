package main

import (
	"log"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/exchanges/bx"
)

// TODO: turn this into tests
func main() {
	clientNoAuth := crypto.Client(bx.New, nil)

	x, err := clientNoAuth.AllOrderBooks()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(x)
}
