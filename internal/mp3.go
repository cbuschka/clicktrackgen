package internal

import (
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

// LoadMp3Samples decodes an MP3 file into a slice of int16 samples.
func LoadMp3Sample(path string) (*Sample, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Initialize the decoder
	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		return nil, err
	}

	// Prepare the buffer for PCM data
	// Note: go-mp3 provides a stream of bytes. 
	// Since it's 16-bit, every 2 bytes = 1 sample.
	var samples []int16
	buf := make([]byte, 4096) // Read in chunks

	for {
		n, err := decoder.Read(buf)
		if n > 0 {
			for i := 0; i < n; i += 2 {
				// Convert 2 bytes (Little Endian) to one int16 sample
				sample := int16(buf[i]) | (int16(buf[i+1]) << 8)
				samples = append(samples, sample)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	// Simple Linear Resampling logic (concept)
	//ratio := float64(decoder.SampleRate()) / 44100.0
	//newLength := int(float64(len(samples)) / ratio)
	//resampled := make([]int16, newLength)

	//for i := 0; i < newLength; i++ {
	//	oldIdx := float64(i) * ratio
	//	resampled[i] = samples[int(oldIdx)] // Simplest "Nearest Neighbor" approach
	//}

	return &Sample{Rate: InternalSampleRate, Data: samples}, nil
}

