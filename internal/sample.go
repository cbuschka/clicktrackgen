package internal

import (
	"errors"
)

const (
	SampleRate = 44100
	BitDepth   = 16
	MaxAmp     = 32767 // Max value for int16
)

var ErrDiffSampleRate = errors.New("different sample rate")

type Sample struct {
	Rate int
	Data []int16
}

func NewSample(rate int, len int) *Sample {
	buf := make([]int16, 0, len)
	return &Sample{Rate: rate, Data: buf}
}

func (s *Sample) Clone() *Sample {
	tmp := make([]int16, len(s.Data))
	copy(tmp, s.Data)
	return &Sample{Rate: s.Rate, Data: tmp}
}

func (s *Sample) MixIn(voice *Sample, offset int, voiceGain float64) error {
	if voice.Rate != s.Rate {
		return ErrDiffSampleRate
	}
	MixAudio(s.Data, voice, offset, voiceGain)
	return nil
}

func (s *Sample) TrimSilence(threshold float64) {
	s.Data = TrimSilence(s.Data, int16(MaxAmp * threshold))
}
