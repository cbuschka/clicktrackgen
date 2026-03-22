package internal

// GenerateCountin creates the 2-measure intro buffer
func (g *Generator) GenerateCountin(samplesPerBeat int, clickAsset *Sample, target *Sample, gain float64) error {
	// Measure 0: Two Half Notes (Beat 1 and Beat 3)
	// Measure 1: Four Quarter Notes (Beat 1, 2, 3, 4)
	
	// Map of beats to trigger: [Measure][Beat]
	timeline := [][]bool{
		{true, false, true, false}, // Measure 0
		{true, true, true, true},   // Measure 1
	}

	for m, beats := range timeline {
		for b, active := range beats {
			if !active {
				continue
			}
			
			offset := (m * 4 * samplesPerBeat) + (b * samplesPerBeat)
			
			err := target.MixIn(clickAsset, offset, gain)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
