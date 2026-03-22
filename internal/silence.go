package internal

import "math"

// TrimSilence removes leading and trailing silence from a PCM slice.
// threshold: the absolute value a sample must exceed to be considered "audio" (e.g., 500).
func TrimSilence(samples []int16, threshold int16) []int16 {
	if len(samples) == 0 {
		return samples
	}

	start := 0
	end := len(samples) - 1

	// 1. Find the first sample that exceeds the threshold (Leading silence)
	for i := 0; i < len(samples); i++ {
		if int16(math.Abs(float64(samples[i]))) > threshold {
			start = i
			break
		}
	}

	// 2. Find the last sample that exceeds the threshold (Trailing silence)
	// This helps prevent "overlapping" if samples are longer than a beat.
	for i := len(samples) - 1; i >= start; i-- {
		if int16(math.Abs(float64(samples[i]))) > threshold {
			end = i
			break
		}
	}

	// Return a slice of the original buffer (Memory efficient)
	return samples[start : end+1]
}
