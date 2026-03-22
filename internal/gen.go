package internal

import (
	"math"
)

const (
	SampleRate = 44100
	BitDepth   = 16
	MaxAmp     = 32767 // Max value for int16
)

type Generator struct {
	BPM          int
	Measures     int
	FileName     string
	CustomSample []int16 // Optional: User-provided WAV data
}

func (g *Generator) Generate() error {
	samplesPerBeat := (SampleRate * 60) / g.BPM
	totalMeasures := g.Measures + 2
	totalSamples := totalMeasures * 4 * samplesPerBeat
	buffer := make([]int16, totalSamples)

	// 1. Determine our "Source Asset"
	var clickAsset []int16
	if len(g.CustomSample) > 0 {
		clickAsset = g.CustomSample
	} else {
		// Fallback to the generated Sine Pulse if no file was provided
		clickAsset = generateSinePulse(1000.0, 0.05)
	}

	// 2. Orchestration Loop
	for m := 0; m < totalMeasures; m++ {
		for b := 0; b < 4; b++ {
			// Skip logic for the first measure of count-in (Beats 3 & 4)
			if m == 0 && b >= 2 {
				continue
			}

			// Calculate the "Address" in the buffer
			offset := (m * 4 * samplesPerBeat) + (b * samplesPerBeat)

			// 3. Use the Mixer instead of a manual copy loop
			// Gain is set to 1.0 for the click to ensure it's the primary clock
			MixAudio(buffer, clickAsset, offset, 1.0)
		}
	}

	return g.writeToWav(buffer)
}

func generateSinePulse(freq float64, duration float64) []int16 {
	numSamples := int(float64(SampleRate) * duration)
	pulse := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / SampleRate
		// Apply a simple linear decay (Envelope) to avoid digital clicking/popping
		envelope := 1.0 - (float64(i) / float64(numSamples))
		value := math.Sin(2 * math.Pi * freq * t)
		pulse[i] = int16(value * MaxAmp * 0.5 * envelope)
	}
	return pulse
}
