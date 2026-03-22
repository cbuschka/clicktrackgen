package internal

import (
	"math"
)

type Generator struct {
	BPM          int
	Measures     int
	ClickTrackFileName     string
	ClueTrackFileName     string
	CustomSample *Sample // Optional: User-provided WAV data
	AccentCustomSample *Sample // Optional: User-provided WAV data
	Clues map[int]string
}

func (g *Generator) Generate() error {
	samplesPerBeat := (SampleRate * 60) / g.BPM
        bufferLen := (g.Measures + 2) * 4 * samplesPerBeat
        buffer := make([]int16, bufferLen)
	clickTrackSample := &Sample{Rate: 44100, Data: buffer}

	err := g.GenerateClickTrack(clickTrackSample)
	if err != nil {
		return err
	}

	err = g.writeToWav(g.ClickTrackFileName, clickTrackSample)
	if err != nil {
		return err
	}

        buffer = make([]int16, bufferLen)
	clueTrackSample := &Sample{Rate: 44100, Data: buffer}
        err = g.GenerateClueStream(samplesPerBeat, clueTrackSample, 1.0)
	if err != nil {
		return err
	}

	err = g.writeToWav(g.ClueTrackFileName, clueTrackSample)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) GenerateClickTrack(target *Sample) error {
	samplesPerBeat := (SampleRate * 60) / g.BPM
	
	// 1. Prepare assets
	var clickAsset *Sample
	var accentClickAsset *Sample
	if g.CustomSample != nil {
		clickAsset = g.CustomSample
	} else {
		clickAsset = generateSinePulse(1000.0, 0.05)
	}
	if g.AccentCustomSample != nil {
		accentClickAsset = g.AccentCustomSample
	} else {
		accentClickAsset = generateSinePulse(1000.0, 0.05)
	}

	// 2. Generate the Count-in "Module"
	err := g.GenerateCountin(samplesPerBeat, clickAsset, target, 1.0)
	if err != nil {
		return err
	}

	samplesForCountIn := samplesPerBeat * 4 * 2 

	// 3. Generate the Main Song "Module"
	songMeasures := g.Measures
	
	for m := 0; m < songMeasures; m++ {
		for b := 0; b < 4; b++ {
			offset := samplesForCountIn + (m * 4 * samplesPerBeat) + (b * samplesPerBeat)

			clickGain := 0.75			
			asset := clickAsset
			if b == 0 {
				clickGain = 1.0
				asset = accentClickAsset
			}
		
			target.MixIn(asset, offset, clickGain)	
		}
	}

	return nil
}

func generateSinePulse(freq float64, duration float64) *Sample {
	numSamples := int(float64(SampleRate) * duration)
	pulse := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / SampleRate
		// Apply a simple linear decay (Envelope) to avoid digital clicking/popping
		envelope := 1.0 - (float64(i) / float64(numSamples))
		value := math.Sin(2 * math.Pi * freq * t)
		pulse[i] = int16(value * MaxAmp * 0.5 * envelope)
	}
	return &Sample{Rate: 44100, Data: pulse}
}
