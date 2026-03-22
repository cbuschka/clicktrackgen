package internal

// MixAudio combines the 'voice' buffer into the 'click' buffer at a specific offset.
// voiceGain should be between 0.0 and 1.0 (e.g., 0.7 for -3dB).
func MixAudio(clickBase []int16, voice []int16, offset int, voiceGain float64) {
	for i := 0; i < len(voice); i++ {
		targetIdx := offset + i
		
		// Boundary check (Safety first)
		if targetIdx >= len(clickBase) {
			break
		}

		// 1. Convert to float64 for high-precision math
		// 2. Apply gain to the voice sample
		// 3. Sum the signals
		mixed := float64(clickBase[targetIdx]) + (float64(voice[i]) * voiceGain)

		// 4. Hard Clipping (Protection against overflow)
		if mixed > 32767 {
			mixed = 32767
		} else if mixed < -32768 {
			mixed = -32768
		}

		// 5. Cast back to int16
		clickBase[targetIdx] = int16(mixed)
	}
}
