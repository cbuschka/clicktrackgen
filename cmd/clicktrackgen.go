package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/cbuschka/clicktrackgen/internal"
)

func main() {
	// 1. Define Flags
	bpm := flag.Int("bpm", 120, "Beats per minute")
	measures := flag.Int("m", 4, "Number of measures (excluding count-in)")
	out := flag.String("o", "click.wav", "Output file name")

	flag.Parse()

	// 2. Validate Input
	if *bpm <= 0 || *measures <= 0 {
		fmt.Println("Error: BPM and Measures must be greater than 0")
		os.Exit(1)
	}

	// 3. Initialize Generator
	gen := &internal.Generator{
		BPM:      *bpm,
		Measures: *measures,
		FileName: *out,
	}

	fmt.Printf("Generating %d BPM click track (%d measures + 2 count-in)...\n", *bpm, *measures)

	// 4. Run the "Job"
	err := gen.Generate()
	if err != nil {
		log.Fatalf("Failed to generate click track: %v", err)
	}

	absPath, _ := filepath.Abs(*out)
	fmt.Printf("Success! File saved to: %s\n", absPath)
}
