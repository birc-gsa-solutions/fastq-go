package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa/fastq"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s fastq\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	err = fastq.ScanFastq(f, func(rec *fastq.FastqRecord) {
		fmt.Println(rec.Name)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
}
