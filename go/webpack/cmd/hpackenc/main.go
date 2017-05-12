// Encodes HTTP headers, specified as alternating names and values on the
// command line, into HPACK. The raw binary is written to stdout.
package main

import (
	"log"
	"os"

	"golang.org/x/net/http2/hpack"
)

func main() {
	encoder := hpack.NewEncoder(os.Stdout)
	for i := 1; i < len(os.Args); i += 2 {
		header := hpack.HeaderField{Name: os.Args[i], Value: os.Args[i+1]}
		if err := encoder.WriteField(header); err != nil {
			log.Fatal(err)
		}
	}
}
