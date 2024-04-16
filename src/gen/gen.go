package gen

import (
	"crypto/rand"
	"log"
)

func Generate(method string) []byte {
	init_key := make([]byte, 64)

	if _, err := rand.Read(init_key); err != nil {
		log.Fatal(err)
	}

	return init_key
}
