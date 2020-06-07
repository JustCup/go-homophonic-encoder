package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func encode(key map[string]characters, text string) (encoded string) {

	buf := strings.Split(text, "")

	for _, chr := range buf {
		if key[chr] == nil {
			randomInt := rand.Intn(900-25) + 25
			key[chr] = append(key[chr], randomInt)
			encoded += fmt.Sprintf("%d ", randomInt)
		} else {
			if len(key[chr]) < 3 {
				randomInt := rand.Intn(900-25) + 25
				key[chr] = append(key[chr], randomInt)
				encoded += fmt.Sprintf("%d ", randomInt)
			} else {
				encoded += fmt.Sprintf("%d ", key[chr][rand.Intn(2)])
			}
		}
	}

	return
}
