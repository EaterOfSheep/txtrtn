package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func reverbMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Feedback: %.2f",feedback))

	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Length: %.2f",reverbLength))


}


func reverbMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){

			case 0: feedback+=0.02
			case 1: reverbLength+=0.02

			}
		case 'h':
			switch(cursorpos){

			case 0: feedback-=0.02
			case 1: reverbLength-=0.02


			}

	}
}
