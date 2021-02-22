package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

type word uint16

func init() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UnixNano())

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
}
