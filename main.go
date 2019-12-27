package main

import (
	"flag"
	"fmt"
	"time"
)

const MB = 1024 * 1024

var ballast []byte

func main() {
	ballastMb := flag.Int("ballast-mb", 0, "ballast size in MB")
	loop := flag.Int("loop", 100, "loop size")
	flag.Parse()

	ballast = make([]byte, *ballastMb*MB)

	start := time.Now()
	for i := 0; i < *loop; i++ {
		tmp := make([]byte, 1*MB)
		if false {
			fmt.Println(tmp)
		}
	}

	end := time.Since(start)
	fmt.Printf("\nelapsed %s\n", end)
}
