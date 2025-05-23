package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s file...\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		fmt.Printf("%s %s\n", checksum(file), file)
	}
}

func checksum(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", sha256.Sum256(b))
}
