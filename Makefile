.PHONY:	build run
build:
	mkdir -p dist/
	go build -o dist/clicktrackgen cmd/clicktrackgen.go

run:
	go run cmd/clicktrackgen.go \
		--accentSample=samples/Perc_Stick_hi.wav \
		--sample=samples/Perc_Stick_lo.wav \
		--clickTrackOut=click.wav \
		--clueTrackOut=clue.wav \
		--combinedTrackOut=combined.wav \
		-m 48 \
		--clues=0:"verse 1",4:"chorus 1",6:"breakdown"
