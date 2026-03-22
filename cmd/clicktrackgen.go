package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"github.com/cbuschka/clicktrackgen/internal"
)

func parseClues(input string) map[int]string {
	clues := make(map[int]string)
	pairs := strings.Split(input, ",")
	for _, p := range pairs {
		parts := strings.Split(p, ":")
		if len(parts) == 2 {
			m, _ := strconv.Atoi(parts[0])
			text := strings.Trim(parts[1], "\"")
			clues[m] = text
		}
	}
	return clues
}

func main() {
	bpm := flag.Int("bpm", 120, "Beats per minute")
	measures := flag.Int("m", 4, "Number of measures")
	out := flag.String("o", "click.wav", "Output file")
	samplePath := flag.String("sample", "", "Path to custom click WAV (optional)")
	accentSamplePath := flag.String("accentSample", "", "Path to custom click WAV (optional)")
	cluesFlag := flag.String("clues", "", "Clues")

	flag.Parse()

	var customData *internal.Sample
	var accentCustomData *internal.Sample
	var err error

	var clues map[int]string
	if *cluesFlag != "" {
		clues = parseClues(*cluesFlag)
	}

	// If the user provided a sample, load it into memory
	if *samplePath != "" {
		customData, err = internal.LoadWavSample(*samplePath)
		if err != nil {
			log.Fatalf("Could not load custom sample: %v", err)
		}
		fmt.Println("Using custom click sample:", *samplePath)
	}

	if *accentSamplePath != "" {
		accentCustomData, err = internal.LoadWavSample(*accentSamplePath)
		if err != nil {
			log.Fatalf("Could not load custom sample: %v", err)
		}
		fmt.Println("Using custom click sample:", *accentSamplePath)
	}

	gen := &internal.Generator{
		BPM:          *bpm,
		Measures:     *measures,
		FileName:     *out,
		CustomSample: customData, // Pass the slice (nil if not loaded)
		AccentCustomSample: accentCustomData, // Pass the slice (nil if not loaded)
		Clues:        clues,
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
