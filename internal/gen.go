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
	AccentCustomSample []int16 // Optional: User-provided WAV data
}

// GenerateCountin creates the 2-measure intro buffer
func (g *Generator) GenerateCountin(samplesPerBeat int, clickAsset []int16) []int16 {
	// 2 measures * 4 beats
	countInSamples := 2 * 4 * samplesPerBeat
	buffer := make([]int16, countInSamples)

	// Measure 0: Two Half Notes (Beat 1 and Beat 3)
	// Measure 1: Four Quarter Notes (Beat 1, 2, 3, 4)
	
	// Map of beats to trigger: [Measure][Beat]
	timeline := [][]bool{
		{true, false, true, false}, // Measure 0
		{true, true, true, true},   // Measure 1
	}

	for m, beats := range timeline {
		for b, active := range beats {
			if !active {
				continue
			}
			
			offset := (m * 4 * samplesPerBeat) + (b * samplesPerBeat)
			
			MixAudio(buffer, clickAsset, offset, 1.0)
		}
	}
	return buffer
}

func (g *Generator) Generate() error {
	samplesPerBeat := (SampleRate * 60) / g.BPM
	
	// 1. Prepare assets
	var clickAsset []int16
	var accentClickAsset []int16
	if len(g.CustomSample) > 0 {
		clickAsset = g.CustomSample
	} else {
		clickAsset = generateSinePulse(1000.0, 0.05)
	}
	if len(g.AccentCustomSample) > 0 {
		accentClickAsset = g.AccentCustomSample
	} else {
		accentClickAsset = generateSinePulse(1000.0, 0.05)
	}

	// 2. Generate the Count-in "Module"
	countInBuf := g.GenerateCountin(samplesPerBeat, clickAsset)

	// 3. Generate the Main Song "Module"
	songMeasures := g.Measures
	songBuf := make([]int16, songMeasures*4*samplesPerBeat)
	
	for m := 0; m < songMeasures; m++ {
		for b := 0; b < 4; b++ {
			offset := (m * 4 * samplesPerBeat) + (b * samplesPerBeat)

			clickGain := 0.75			
			asset := clickAsset
			if b == 0 {
				clickGain = 1.0
				asset = accentClickAsset
			}
			
			MixAudio(songBuf, asset, offset, clickGain)
		}
	}

	// 4. Concatenate/Combine (The "Final Build")
	// We create a master buffer and copy the parts in
	finalBuffer := append(countInBuf, songBuf...)

	return g.writeToWav(finalBuffer)
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
