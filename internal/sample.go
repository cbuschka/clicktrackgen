package internal

const (
	SampleRate = 44100
	BitDepth   = 16
	MaxAmp     = 32767 // Max value for int16
)

type Sample struct {
	Rate int
	Data []int16
}

func NewSample(rate int, len int) *Sample {
	buf := make([]int16, 0, len)
	return &Sample{Rate: rate, Data: buf}
}

func (s *Sample) MixIn(voice *Sample, offset int, voiceGain float64) {
	MixAudio(s.Data, voice, offset, voiceGain)
}

