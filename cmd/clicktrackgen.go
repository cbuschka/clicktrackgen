package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"github.com/cbuschka/clicktrackgen/internal"
)

func parseClues(input string) ([]internal.Clue, error) {
	clues := make([]internal.Clue, 0, 0)
	pairs := strings.Split(input, ",")
	for _, p := range pairs {
		parts := strings.Split(p, ":")
		if len(parts) == 2 {
			m, _ := strconv.Atoi(parts[0])
			text := strings.Trim(parts[1], "\"")
			clues = append(clues, internal.Clue{Bar: m, Name: text})
		}
	}
	return clues, nil
}

func main() {
	var err error

	bpm := flag.Int("bpm", 120, "Beats per minute")
	bars := flag.Int("bars", 4, "Number of bars")
	beatsPerBar := flag.Int("beatsPerBar", 4, "Number of beats per bar")
	songTrackIn := flag.String("songTrackIn", "song.wav", "Input file")
	clickTrackOut := flag.String("clickTrackOut", "click.wav", "Output file")
	clueTrackOut := flag.String("clueTrackOut", "clue.wav", "Output file")
	combinedTrackOut := flag.String("combinedTrackOut", "combined.wav", "Output file")
	samplePath := flag.String("sample", "", "Path to custom click WAV (optional)")
	accentSamplePath := flag.String("accentSample", "", "Path to custom click WAV (optional)")
	cluesFlag := flag.String("clues", "", "Clues")

	flag.Parse()

	var customData *internal.Sample
	var accentCustomData *internal.Sample

	var clues []internal.Clue
	if *cluesFlag != "" {
		clues, err = parseClues(*cluesFlag)
		if err != nil {
			log.Fatalf("Could not parse clues: %v", err)
		}
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
		Bars:     *bars,
		BeatsPerBar: *beatsPerBar,
		SongTrackFileName: *songTrackIn,
		ClickTrackFileName:     *clickTrackOut,
		ClueTrackFileName:     *clueTrackOut,
		CombinedTrackFileName:     *combinedTrackOut,
		CustomSample: customData, // Pass the slice (nil if not loaded)
		AccentCustomSample: accentCustomData, // Pass the slice (nil if not loaded)
		Clues:        clues,
	}

	fmt.Printf("Generating %d BPM click track (%d bars + 2 count-in)...\n", *bpm, *bars)

	// 4. Run the "Job"
	err = gen.Generate()
	if err != nil {
		log.Fatalf("Failed to generate click track: %v", err)
	}
}
