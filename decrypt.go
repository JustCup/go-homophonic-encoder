package main

import (
	"strconv"
	"strings"
)

func decode(key map[string]characters, text string) (decoded string) {

	buf := strings.Split(text, " ")

	for _, chr := range buf {
		for k, v := range key {
			for _, val := range v {
				if strconv.Itoa(val) == chr {
					decoded += k
					continue
				}
			}
		}
	}

	return
}
