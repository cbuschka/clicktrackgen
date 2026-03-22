package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"github.com/cbuschka/clicktrackgen/internal"
)

func main() {
	bpm := flag.Int("bpm", 120, "Beats per minute")
	measures := flag.Int("m", 4, "Number of measures")
	out := flag.String("o", "click.wav", "Output file")
	samplePath := flag.String("sample", "", "Path to custom click WAV (optional)")

	flag.Parse()

	var customData []int16
	var err error

	// If the user provided a sample, load it into memory
	if *samplePath != "" {
		customData, err = internal.LoadWavSamples(*samplePath)
		if err != nil {
			log.Fatalf("Could not load custom sample: %v", err)
		}
		fmt.Println("Using custom click sample:", *samplePath)
	}

	gen := &internal.Generator{
		BPM:          *bpm,
		Measures:     *measures,
		FileName:     *out,
		CustomSample: customData, // Pass the slice (nil if not loaded)
	}

	fmt.Printf("Generating %d BPM click track (%d measures + 2 count-in)...\n", *bpm, *measures)

	// 4. Run the "Job"
	err = gen.Generate()
	if err != nil {
		log.Fatalf("Failed to generate click track: %v", err)
	}

	absPath, _ := filepath.Abs(*out)
	fmt.Printf("Success! File saved to: %s\n", absPath)
}
