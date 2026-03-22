package internal

import (
	"encoding/binary"
	"os"
)

// WriteWav wraps the raw PCM data with the necessary 44-byte header
func (g *Generator) writeToWav(buffer []int16) error {
	f, err := os.Create(g.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

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
	binary.Write(f, binary.LittleEndian, uint32(SampleRate))
	// Byte Rate (SampleRate * NumChannels * BitsPerSample/8)
	binary.Write(f, binary.LittleEndian, uint32(SampleRate*1*2))
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
