package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bitfield/shift"
)

func main() {
	crib := flag.String("crib", "", "crib text")
	flag.Parse()
	if *crib == "" {
		fmt.Fprintln(os.Stderr, "Please specify a crib text with -crib")
		os.Exit(1)
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	key, err := shift.Crack(data, []byte(*crib))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Stdout.Write(shift.Decipher(data, key))
}
