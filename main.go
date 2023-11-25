package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func run() error {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	pages := readPages(string(bytes))

	fmt.Printf("total: %d, w: %d, h: %d\n", len(pages.ps), pages.width, pages.height)
	for _, p := range pages.ps {
		p.print()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
