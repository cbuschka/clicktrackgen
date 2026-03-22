Background
===========
A click track is an audio file that contains a range of clicks of equal distance corresponding to the velocity
of a song measured in bpm (beats per minutes). The most basic version is a click sound per beat.

A click track starts with counting in, two measures. The first one half of the beats, the second all beats like 1, 2,
, 1, 2, 3, 4.

The click track is a mono audio file, mostly wav or raw audio, 44.1 khz, signed 16bit audio. The sound can be a sinus
style click or ping sound.

A click track can be accompanied by a second mono audio stream which contains clues for the musician spoken by a human
voice. In case of a drummer something like "verse 1" or "bridge" or "breakdown" with a subsequent countdown when it starts
"4 3 2 1".

The click track generator
=========================
The click generator generates an audio file like specified as above for a given bpm input, a number of measures,
 and a file name to write the audio data to. The click track is generated the 
following: First the click sample is generated from a sinus pulse as a audio data.
An empty audio stream is generated, the countin is mixed into it as well as the clicks requested. Count in is not part of the
number of measures. So the audio is two mesaures longer than requested.

Technical information
=====================
It is written in go and utilizes library. 

Via flags the output file name is read from the command line and also the bpm and number of measures.

Then there is a cmd/clicktrackgen that evaluates the command line flags. Unter internal/ there is a command, that can be
configured with the flags. The output is written afterward to an file.
