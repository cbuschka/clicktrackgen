package idorup1

import "io"

// in lists/session.json:
// {"id":"d7aa13cc-3002-4152-8ce3-980e06b7d2f4","filePath":"/home/conni/Music/Idoru-p1/Setlist1.idoru","name":"Unnamed Session","deviceImport":false}
type SessionJsonFile struct {
	Id           string `json:"id"`
	FilePath     string `json:"filePath"`
	Name         string `json:"name"`
	DeviceImport bool   `json:"deviceImport"`
}

func ReadSessionJsonFile(r io.Reader) (*SessionJsonFile, error) {
	return nil, nil
}

func (f *SessionJsonFile) Write(w io.Writer) error {
	return nil
}

// in lists/Test/setlist.json:
// {"id":"7e25c752-7188-4e12-9d68-889a631824f5","songs":["2910a252-76f6-496c-a74e-2ccd3516ce8e"]}
type SetlistJsonFile struct {
	Id    string   `json:"id"`
	Songs []string `json:"songs"`
}

func ReadSetlistJsonFile(r io.Reader) (*SetlistJsonFile, error) {
	return nil, nil
}

func (f *SetlistJsonFile) Write(w io.Writer) error {
	return nil
}

// in lists/Test/Test.txt:
// SetList file
//
// Global sets
// StereoLinks-
// HeadPhone- 95
// Output1- 95
// Output2- 95
// Output3- 95
// Output4- 95
// Output5- 95
// Output6- 95
//
// Songs
// "First Song"
type SetlistTxtFile struct {
}

func ReadSetlistTxtFile(r io.Reader) (*SetlistTxtFile, error) {
	return nil, nil
}

func (f *SetlistTxtFile) Write(w io.Writer) error {
	return nil
}

// in lists/Test/First Song/song.json:
// {"id":"2910a252-76f6-496c-a74e-2ccd3516ce8e"}
type SongJsonFile struct {
	Id string
}

func ReadSongJsonFile(r io.Reader) (*SongJsonFile, error) {
	return nil, nil
}

func (f *SongJsonFile) Write(w io.Writer) error {
	return nil
}

// in lists/Test/First Song/fileMap.json
// {"b8839551-d5f1-44ec-bef2-77a6ee595775":"Siren-441_s16le.wav"}
type FileMapJson map[string]string

func ReadFileMapJson(r io.Reader) (*FileMapJson, error) {
	return nil, nil
}

func (f *FileMapJson) Write(w io.Writer) error {
	return nil
}

// in lists/Test/First Song/First Song.txt:
//
// Global sets
// Level- 0
// BPM- 120
// AtEnd- QueueNext
//
// Input Files
// F1- "F1" "Siren-441_s16le"
// F2- "F2" ""
// F3- "F3" ""
// F4- "F4" ""
// F5- "F5" ""
// F6- "F6" ""
// MIDI-
//
// HeadPhone
// IN1- F1 95
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output1
// IN1- F1 95
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output2
// IN1- F1 95 MUTE
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output3
// IN1- F1 95 MUTE
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output4
// IN1- F1 95 MUTE
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output5
// IN1- F1 95 MUTE
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
//
// Output6
// IN1- F1 95 MUTE
// IN2- F2 95 MUTE
// IN3- F3 95 MUTE
// IN4- F4 95 MUTE
// IN5- F5 95 MUTE
// IN6- F6 95 MUTE
// IN7- AN 95 MUTE
type SongTxtFile struct {
}

func ReadSongTxtFile(r io.Reader) (*SongTxtFile, error) {
	return nil, nil
}

func (f *SongTxtFile) Write(w io.Writer) error {
	return nil
}
