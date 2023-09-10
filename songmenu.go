package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var editsong int = 0

func songMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Song %d: %s",editsong,songs[editsong].name))
	drawText(*s, 5, 4, 50, 4, style, "Play")
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Looping: %t",songs[editsong].looping))

	drawText(*s, 5, 9, 50,9, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(songs[editsong].clip.wave))/44100.0, intPercentage((int)(songs[editsong].clip.age),len(songs[editsong].clip.wave))))



}




func songMenuControl(r rune){

	switch(r){


		case 'K':
			editsong--
		case 'J':
			editsong++


		case 'l':
			switch(cursorpos){
				case 0:
					editsong++
				case 1:
					songs[editsong].play()
				case 3:
					songs[editsong].looping = !songs[editsong].looping


			}
		case 'h':
			switch(cursorpos){
				case 0:
					editsong--


			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
