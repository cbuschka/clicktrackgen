package internal

import (
	"math"
)

const (
	SampleRate = 44100
	BitDepth   = 16
	MaxAmp     = 32767 // Max value for int16
)

// Generator handles the audio buffer construction
type Generator struct {
	BPM      int
	Measures int
	FileName string
}

// Generate creates the raw PCM data and writes to file
func (g *Generator) Generate() error {
	samplesPerBeat := (SampleRate * 60) / g.BPM
	totalMeasures := g.Measures + 2 // Including 2-measure count-in
	totalSamples := totalMeasures * 4 * samplesPerBeat
	buffer := make([]int16, totalSamples)

	// 2. Map the Clicks to the Buffer
	for m := 0; m < totalMeasures; m++ {
		for b := 0; b < 4; b++ {
			currentIndex := (m * 4 * samplesPerBeat) + (b * samplesPerBeat)

			// Logic for Count-in (Measure 0: only beats 1 & 2. Measure 1: all beats)
			if m == 0 && b >= 2 {
				continue // Silence for beats 3 & 4 of first measure
			}

			// Logic for the Accent (Beat 1 of the measure)
			freq := 1000.0 // Standard "beep"
			if b == 0 {
				freq = 1500.0 // Higher pitch for the "One"
			}

			tickSamples := generateSinePulse(freq, 0.05)

			// Write the tick into the buffer
			for i := 0; i < len(tickSamples) && (currentIndex+i) < len(buffer); i++ {
				buffer[currentIndex+i] = tickSamples[i]
			}
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
