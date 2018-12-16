package main

import (
	"math/rand"
	"time"

	"github.com/andersnormal/outlaw/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
