package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var editsample int = 0

func sampleMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Sample %d: %s",editsample,samples[editsample].name))

	drawText(*s, 5, 4, 50, 4, style, "Fire!")

	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(samples[editsample].clip.wave))/44100.0, intPercentage(samples[editsample].clip.age,len(samples[editsample].clip.wave))))



}




func sampleMenuControl(r rune){

	switch(r){


		case 'K':
			editsample--
		case 'J':
			editsample++


		case 'l':
			switch(cursorpos){
				case 0:
					editsample++
				case 1:
					samples[editsample].play()


			}
		case 'h':
			switch(cursorpos){
				case 0:
					editsample--

			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
