package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mathrand "math/rand"
	"time"
)

func GenerateRandomOwner() string {
	seed1 := GenerateCryptRandomNumber()
	seed2 := time.Now().UnixNano()

	owner := sha256.Sum256([]byte(fmt.Sprintf("%s%d", seed1, seed2)))
	// return string(owner[:])

	return hex.EncodeToString(owner[:])
}

func GenerateCryptRandomNumber() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func generate_random_num(maxnum uint32) uint64 {
	for {
		seed2 := time.Now().UnixNano()
		mathrand.Seed(seed2)
		amt := mathrand.Uint64() % uint64(maxnum)
		if amt > 0 {
			return amt
		}
	}
}

func generate_random_num_in_range(min, max int) int {
	if min < 1 {
		min = 1
	}
	if max <= min {
		max = min + 2
	}

	for {
		seed2 := time.Now().UnixNano()
		// mathrand.Seed(seed2)
		mathrand.New(mathrand.NewSource(seed2))
		amt := mathrand.Uint64() % uint64(max)
		if amt >= uint64(min) {
			return int(amt)
		}
	}
}
