package internal

import (
	"bufio"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	shineMp3 "github.com/braheezy/shine-mp3/pkg/mp3"
)

func loadMp3Sample(path string) (*Sample, error) {
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

func writeMp3Sample(path string, sample *Sample) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	out := bufio.NewWriter(f)
	defer out.Flush()

	mp3Encoder := shineMp3.NewEncoder(sample.Rate, 2)

	data := make([]int16, 0, len(sample.Data)*2)
	for index := range sample.Data {
		data = append(data, sample.Data[index], sample.Data[index])
	}

	err = mp3Encoder.Write(out, data)
	if err != nil {
		return err
	}

	return nil
}
