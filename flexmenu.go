package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func flexMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Flexing: %v",flexing))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Flex Type: %s",flexTypeName()))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Flex Scale: %.2f",flexscale))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Frequency: %v/4sbar",flexFreq))

}


func flexMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){

			case 0: flexing=!flexing
			case 1: flextype++
			case 2: flexscale+=0.05
			case 3: flexFreq++

			}
		case 'h':
			switch(cursorpos){
				case 1: flextype--
				case 2: flexscale-=0.05
				case 3: flexFreq--


			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){


			}
	}
}
