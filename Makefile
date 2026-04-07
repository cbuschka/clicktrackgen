.PHONY:	build run
build:
	mkdir -p dist/
	go build -o dist/clicktrackgen cmd/clicktrackgen.go

run:	build
	./dist/clicktrackgen \
		--accentSample=samples/Perc_Stick_hi.wav \
		--sample=samples/Perc_Stick_lo.wav \
		--songTrackIn=song.wav \
		--clickTrackOut=click.mp3 \
		--clueTrackOut=clue.mp3 \
		--combinedTrackOut=combined.mp3 \
		--allTrackOut=all.mp3 \
		--bpm 150 \
		--bars 156 \
		--clues=0:"intro",8:"verse 1",24:"pre chorus",32:"chorus 1",48:"interlude",56:"verse 2",72:"pre chorus 2",80:"chorus 2",96:"interlude 2",104:"bridge",112:"chorus 3",128:"chorus 4",148:"trash can ending"
