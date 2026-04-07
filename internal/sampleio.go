package internal

import (
	"strings"
	"errors"
)

var ErrSampleFormatNotSupported = errors.New("only wav and mp3 supported")

func ReadSample(path string) (*Sample, error) {
	if strings.HasSuffix(path,".wav") {
		return loadWavSample(path)
	} else if strings.HasSuffix(path,".mp3") {
		return loadMp3Sample(path)
	}

	return nil, ErrSampleFormatNotSupported
}

func WriteSample(path string, sample *Sample) error {
	if strings.HasSuffix(path,".wav") {
		return writeWavSample(path, sample)
	} else if strings.HasSuffix(path,".mp3") {
		return writeMp3Sample(path, sample)
	}

	return ErrSampleFormatNotSupported
}
