// +build ignore

// Read blocks out of a single file to time the seeking code
package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	// Flags
	iterations   = flag.Int("n", 25, "Iterations to try")
	maxBlockSize = flag.Int("b", 1024*1024, "Max block size to read")
	randSeed     = flag.Int64("seed", 1, "Seed for the random number generator")
)

func randomSeekTest(size int64, in *os.File, name string) {
	start := rand.Int63n(size)
	blockSize := rand.Intn(*maxBlockSize)
	if int64(blockSize) > size-start {
		blockSize = int(size - start)
	}
	log.Printf("Reading %d from %d", blockSize, start)

	_, err := in.Seek(start, io.SeekStart)
	if err != nil {
		log.Printf("Seek failed on %q: %v", name, err)
	}

	buf := make([]byte, blockSize)
	_, err = io.ReadFull(in, buf)
	if err != nil {
		log.Printf("Read failed on %q: %v", name, err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Printf("Require 1 file as argument")
	}
	rand.Seed(*randSeed)

	name := args[0]
	in, err := os.Open(name)
	if err != nil {
		log.Printf("Couldn't open %q: %v", name, err)
	}

	fi, err := in.Stat()
	if err != nil {
		log.Printf("Couldn't stat %q: %v", name, err)
	}

	start := time.Now()
	for i := 0; i < *iterations; i++ {
		randomSeekTest(fi.Size(), in, name)
	}
	dt := time.Since(start)
	log.Printf("That took %v for %d iterations, %v per iteration", dt, *iterations, dt/time.Duration(*iterations))

	err = in.Close()
	if err != nil {
		log.Printf("Error closing %q: %v", name, err)
	}
}
