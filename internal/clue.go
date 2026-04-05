package internal

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"net/url"
	"os"
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
	escapedText := url.QueryEscape(text)
	filename := filepath.Join(speech.Folder, escapedText + ".mp3")

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		log.Printf("generating speech file for text='%s'", text)
		voiceFile, err := speech.CreateSpeechFile(text, escapedText)
		if err != nil {
			return nil, err
		}
		filename = voiceFile
	} else {
		log.Printf("speech file for text='%s' already exists", text)
	}

	voiceSample, err := LoadMp3Sample(filename)
	if err != nil {
		return nil, err
	}

	voiceSample.TrimSilence(0.08)

	return voiceSample, nil
}

// GenerateClueStream creates a dedicated mono track for voice cues
func (g *Generator) GenerateClueStream(samplesPerBeat int, target *Sample, gain float64) error {
	tmpDir, err := os.MkdirTemp("", "clicktrackgen")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)
	speech := htgotts.Speech{Folder: tmpDir, Language: voices.English}

	for targetBar, text := range g.Clues {
		// Calculate internal indices. 
		// Remember: We have a 2-measure count-in at the start of the buffer.
		actualTargetBar := targetBar + 2
		
		// 1. Place the "1 2 3 4" countdown in the measure BEFORE the target
		countdownBar := actualTargetBar - 1
		labelBar := actualTargetBar - 1
		if countdownBar >= 0 && labelBar >= 0 {
			for b := 0; b < 4; b++ {
				var err error
				countText := fmt.Sprintf("%d", 4-b)
				// Fetch/Generate TTS for "4", "3", etc.
				voiceSample, err := newSpeechSample(speech, countText)
				if err != nil {
					return err
				}
				
				offset := (countdownBar * 4 * samplesPerBeat) + (b * samplesPerBeat)
				err = target.MixIn(voiceSample, offset, gain)
				if err != nil {
					return err
				}
			}

			// 2. Place the Label ("Verse 1") in the measure BEFORE the countdown
			voiceSample, err := newSpeechSample(speech, text)
			if err != nil {
				return err
			}
			
			offset := ((labelBar * 4 -1) * samplesPerBeat) - len(voiceSample.Data)
			if offset >= 0 {
				err := target.MixIn(voiceSample, offset, gain)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
