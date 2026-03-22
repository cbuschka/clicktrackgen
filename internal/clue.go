package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	htgotts "github.com/hegedustibor/htgo-tts" // Example TTS lib
	voices "github.com/hegedustibor/htgo-tts/voices"
)

type speechHandlerWrapper struct {
	Func func(fileName string) error
}

func (s *speechHandlerWrapper) Play(fileName string) error {
	return s.Func(fileName)
}

func speechHandler(f func(string) error) *speechHandlerWrapper {
	return &speechHandlerWrapper{Func: f}
}

func newSpeechSample(speech htgotts.Speech, text string) (*Sample, error) {
	tmpfile, err := ioutil.TempFile("", "")
	tmpfileName := filepath.Base(tmpfile.Name())

	voiceFile, err := speech.CreateSpeechFile(text, tmpfileName)
	if err != nil {
		return nil, err
	}
	voiceSample, err := LoadMp3Sample(voiceFile)
	if err != nil {
		return nil, err
	}

	return voiceSample,nil
}

// GenerateClueStream creates a dedicated mono track for voice cues
func (g *Generator) GenerateClueStream(samplesPerBeat int, target *Sample, gain float64) error {
	speech := htgotts.Speech{Folder: "audio_cache", Language: voices.English}

	for targetMeasure, text := range g.Clues {
		// Calculate internal indices. 
		// Remember: We have a 2-measure count-in at the start of the buffer.
		actualTargetMeasure := targetMeasure + 2
		
		// 1. Place the "4 3 2 1" countdown in the measure BEFORE the target
		countdownMeasure := actualTargetMeasure - 1
		if countdownMeasure >= 0 {
			for b := 0; b < 4; b++ {
				countText := fmt.Sprintf("%d", 4-b)
				// Fetch/Generate TTS for "4", "3", etc.
				voiceSample, err := newSpeechSample(speech, countText)
				if err != nil {
					return err
				}
				
				offset := ((countdownMeasure * 4 - 1) * samplesPerBeat) + (b * samplesPerBeat)
				target.MixIn(voiceSample, offset, gain)
			}
		}

		// 2. Place the Label ("Verse 1") in the measure BEFORE the countdown
		labelMeasure := actualTargetMeasure - 1
		if labelMeasure >= 0 {
			voiceSample, err := newSpeechSample(speech, text)
			if err != nil {
				return err
			}
			
			offset := (labelMeasure * 4 * samplesPerBeat) - samplesPerBeat - len(voiceSample.Data)
			if offset >= 0 {
				target.MixIn(voiceSample, offset, gain)
			}
		}
	}
	return nil
}
