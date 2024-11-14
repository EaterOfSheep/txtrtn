package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var alphasong = 0;
var betasong = 1;


func djMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Deck A: %d: %s",alphasong,songs[alphasong].name))
	drawText(*s, 5, 4, 50, 4, style, alphaPlayButtonText())
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Volume: %.2f%%",songs[alphasong].clip.volume*100))
	drawText(*s, 5, 6, 50,6, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(songs[alphasong].clip.wave))/44100.0, intPercentage((int)(songs[alphasong].clip.age),len(songs[alphasong].clip.wave))))


	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Deck B: %d: %s",betasong,songs[betasong].name))
	drawText(*s, 5, 9, 50, 9, style, betaPlayButtonText())
	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Volume: %.2f%%",songs[betasong].clip.volume*100))
	drawText(*s, 5, 11, 50,11, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(songs[betasong].clip.wave))/44100.0, intPercentage((int)(songs[betasong].clip.age),len(songs[betasong].clip.wave))))



}

func alphaPlayButtonText() string{

	if(songs[alphasong].playing){return "Playing"}
	if(songs[alphasong].queued){return "Queued..."}
	return "Play"
}

func betaPlayButtonText() string{

	if(songs[betasong].playing){return "Playing"}
	if(songs[betasong].queued){return "Queued..."}
	return "Play"
}







func djMenuControl(r rune){

	switch(r){


		case 'K':
			alphasong--
		case 'J':
			alphasong++


		case 'l':
			switch(cursorpos){
				case 0:
					alphasong++
				case 1:
					songs[alphasong].queue()
				case 2:
					if(songs[alphasong].clip.volume<4){songs[alphasong].clip.volume+=0.05}else{songs[alphasong].clip.volume=4}

				case 5:
					betasong++
				case 6:
					songs[betasong].queue()
				case 7:
					if(songs[betasong].clip.volume<4){songs[betasong].clip.volume+=0.05}else{songs[betasong].clip.volume=4}



			}
		case 'h':
			switch(cursorpos){
				case 0:
					alphasong--
				case 2:
					if(songs[alphasong].clip.volume>0){songs[alphasong].clip.volume-=0.05}else{songs[alphasong].clip.volume=0}

				case 5:
					betasong--

				case 7:
					if(songs[betasong].clip.volume>0){songs[betasong].clip.volume-=0.05}else{songs[betasong].clip.volume=0}




			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
