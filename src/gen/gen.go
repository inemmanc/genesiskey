package gen

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func Generate(method string) string {
	init_key := make([]byte, 64)

	if _, err := rand.Read(init_key); err != nil {
		log.Fatal(err)
	}

	switch method {
	case "std":
		return base64.StdEncoding.EncodeToString(init_key)
	case "raw":
		return base64.RawStdEncoding.EncodeToString(init_key)
	}
	
	return ""
}
