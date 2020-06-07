package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type characters []int

func main() {
	rand.Seed(time.Now().UnixNano())

	Key := make(map[string]characters)
	encoded := encode(Key, "Testing encode system")

	fmt.Println("Encoded string:", encoded)
	fmt.Println("Generated key:", Key)

	j, _ := json.Marshal(Key)
	ioutil.WriteFile("key.txt", j, 0644)
	ioutil.WriteFile("encoded.txt", []byte(encoded), 0644)

	encodedTXT, _ := ioutil.ReadFile("encoded.txt")
	keyTXT, _ := ioutil.ReadFile("key.txt")

	var key map[string]characters
	json.Unmarshal(keyTXT, &key)

	decoded := decode(key, string(encodedTXT))

	fmt.Println("Decoded string:", decoded)
	ioutil.WriteFile("decoded.txt", []byte(decoded), 0644)
}
