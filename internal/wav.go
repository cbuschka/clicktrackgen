package internal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func writeWavSample(fileName string, sample *Sample) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	f := bufio.NewWriter(file)
	defer f.Flush()

	buffer := sample.Data
	numSamples := len(buffer)
	dataSize := numSamples * 2 // 16-bit = 2 bytes per sample

	// WAV Header Construction
	// -----------------------
	// Chunk ID "RIFF"
	f.Write([]byte("RIFF"))
	// Chunk Size (36 + dataSize)
	binary.Write(f, binary.LittleEndian, uint32(36+dataSize))
	// Format "WAVE"
	f.Write([]byte("WAVE"))
	// Subchunk1 ID "fmt "
	f.Write([]byte("fmt "))
	// Subchunk1 Size (16 for PCM)
	binary.Write(f, binary.LittleEndian, uint32(16))
	// Audio Format (1 for PCM)
	binary.Write(f, binary.LittleEndian, uint16(1))
	// Num Channels (1 for Mono)
	binary.Write(f, binary.LittleEndian, uint16(1))
	// Sample Rate (44100)
	binary.Write(f, binary.LittleEndian, uint32(InternalSampleRate))
	// Byte Rate (InternalSampleRate * NumChannels * BitsPerSample/8)
	binary.Write(f, binary.LittleEndian, uint32(InternalSampleRate*1*2))
	// Block Align (NumChannels * BitsPerSample/8)
	binary.Write(f, binary.LittleEndian, uint16(2))
	// Bits Per Sample (16)
	binary.Write(f, binary.LittleEndian, uint16(16))
	// Subchunk2 ID "data"
	f.Write([]byte("data"))
	// Subchunk2 Size
	binary.Write(f, binary.LittleEndian, uint32(dataSize))

	// Write the actual audio Payload
	for _, sample := range buffer {
		binary.Write(f, binary.LittleEndian, sample)
	}

	return nil
}

func loadWavSample(path string) (*Sample, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 1. Skip the 44-byte header (Assuming you know it's 44.1k/16bit/Mono)
	// In a production tool, you'd parse this to ensure compatibility.
	header := make([]byte, 44)
	if _, err := f.Read(header); err != nil {
		return nil, err
	}

	// Validate it's actually a WAVE file
	if string(header[0:4]) != "RIFF" || string(header[8:12]) != "WAVE" {
		return nil, fmt.Errorf("file %s is not a valid WAV file", path)
	}

	// 2. Read the rest of the file as int16 samples
	var samples []int16
	for {
		var sample int16
		// WAV data is Little Endian
		err := binary.Read(f, binary.LittleEndian, &sample)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		samples = append(samples, sample)
	}

	return &Sample{Rate: InternalSampleRate, Data: samples}, nil
}
