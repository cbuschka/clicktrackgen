package idorup1

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadSessionJsonFile(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected SessionJsonFile
	}{
		{
			name:     "Standard 120BPM Measure 1",
			input:    []byte{0x78, 0x00, 0x01, 0x00, 0x00}, // 120 (0x78), 1, false
			expected: SessionJsonFile{BPM: 120, Measure: 1, IsAccent: false},
		},
		{
			name:     "Accented 144BPM Measure 16",
			input:    []byte{0x90, 0x00, 0x10, 0x00, 0x01}, // 144 (0x90), 16, true
			expected: SessionJsonFile{BPM: 144, Measure: 16, IsAccent: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.input)
			got, err := ReadMetadata(reader)
			if err != nil {
				t.Fatalf("ReadMetadata() unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ReadMetadata() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWriteMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    ClickMetadata
		expected []byte
	}{
		{
			name:     "Write 120BPM",
			input:    ClickMetadata{BPM: 120, Measure: 1, IsAccent: false},
			expected: []byte{0x78, 0x00, 0x01, 0x00, 0x00},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteMetadata(tt.input)
			if err != nil {
				t.Fatalf("WriteMetadata() unexpected error: %v", err)
			}
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("WriteMetadata() = %v, want %v", got, tt.expected)
			}
		})
	}
}
