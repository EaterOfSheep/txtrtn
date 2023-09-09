package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func audioMenuDisplay(s *tcell.Screen, style tcell.Style){


	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Drums Volume: %.2f",drumsVolume))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Samples Volume: %.2f",samplesVolume))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Songs Volume: %.2f",songsVolume))

	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Reverb Feedback: %.2f",feedback))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Reverb Buffer Length: %.2f",reverbLength))


}


func audioMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){

			case 0: drumsVolume+=0.02
			case 1: samplesVolume+=0.02
			case 2: songsVolume+=0.02

			case 4: feedback+=0.02
			case 5: reverbLength+=0.02

			}
		case 'h':
			switch(cursorpos){

			case 0: drumsVolume-=0.02
			case 1: samplesVolume-=0.02
			case 2: songsVolume-=0.02

			case 4: feedback-=0.02
			case 5: reverbLength-=0.02


			}

	}
}
