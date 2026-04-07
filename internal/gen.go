package internal

import (
	"math"
)

type Generator struct {
	BPM          int
	Bars int
	BeatsPerBar int
	CountInBars int
	SongTrackFileName     string
	ClickTrackFileName     string
	ClueTrackFileName     string
	CombinedTrackFileName     string
	CustomSample *Sample // Optional: User-provided WAV data
	AccentCustomSample *Sample // Optional: User-provided WAV data
	Clues []Clue
}

type Clue struct {
	Name string
	Bar int
}

func (g *Generator) Generate() error {
	samplesPerBeat := (InternalSampleRate * 60) / g.BPM
        bufferLen := (g.Bars + g.CountInBars) * g.BeatsPerBar * samplesPerBeat
        buffer := make([]int16, bufferLen)
	clickTrackSample := &Sample{Rate: InternalSampleRate, Data: buffer}

	err := g.GenerateClickTrack(clickTrackSample)
	if err != nil {
		return err
	}

	err = WriteSample(g.ClickTrackFileName, clickTrackSample)
	if err != nil {
		return err
	}

        buffer = make([]int16, bufferLen)
	clueTrackSample := &Sample{Rate: InternalSampleRate, Data: buffer}
        err = g.GenerateClueStream(samplesPerBeat, clueTrackSample, 1.0)
	if err != nil {
		return err
	}

	err = WriteSample(g.ClueTrackFileName, clueTrackSample)
	if err != nil {
		return err
	}


	if g.CombinedTrackFileName != "" {
		combinedTrackSample := clickTrackSample.Clone()
		if g.SongTrackFileName != "" {
			songTrackSample, err := ReadSample(g.SongTrackFileName)
			if err != nil {
				return err
			}

			songTrackSample.TrimSilence(0.1)

			err = combinedTrackSample.MixIn(songTrackSample, g.CountInBars * g.BeatsPerBar * samplesPerBeat, 1.0)
			if err != nil {
				return err
			}
		}

		err = combinedTrackSample.MixIn(clueTrackSample, 0, 1.0)
		if err != nil {
			return err
		}

		err = WriteSample(g.CombinedTrackFileName, combinedTrackSample)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) GenerateClickTrack(target *Sample) error {
	samplesPerBeat := (InternalSampleRate * 60) / g.BPM
	
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

	samplesForCountIn := samplesPerBeat * g.BeatsPerBar * g.CountInBars

	// 3. Generate the Main Song "Module"
	songBars := g.Bars
	
	for m := 0; m < songBars; m++ {
		for b := 0; b < g.BeatsPerBar; b++ {
			offset := samplesForCountIn + (m * g.BeatsPerBar * samplesPerBeat) + (b * samplesPerBeat)

			clickGain := 0.75			
			asset := clickAsset
			if b == 0 {
				clickGain = 1.0
				asset = accentClickAsset
			}
		
			err := target.MixIn(asset, offset, clickGain)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generateSinePulse(freq float64, duration float64) *Sample {
	numSamples := int(float64(InternalSampleRate) * duration)
	pulse := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / InternalSampleRate
		// Apply a simple linear decay (Envelope) to avoid digital clicking/popping
		envelope := 1.0 - (float64(i) / float64(numSamples))
		value := math.Sin(2 * math.Pi * freq * t)
		pulse[i] = int16(value * MaxAmp * 0.5 * envelope)
	}
	return &Sample{Rate: InternalSampleRate, Data: pulse}
}
