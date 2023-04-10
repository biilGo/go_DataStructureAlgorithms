package main

import (
	"fmt"
	"xxhash"
	// "github.com/OneOfOne/xxhash"
)

// 将一个键进行Hash
func XXHahs(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

func main() {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		fmt.Printf("xxhash('%s')=%d\n", key, XXHahs([]byte(key)))
	}
}
