package main

import (
	"io"
	"log"
	"os"
)

func run() error {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	pages, err := readPages(string(bytes))
	if err != nil {
		return err
	}

	return start(pages)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
