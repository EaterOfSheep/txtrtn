package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func flexMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Flex Scale: %v",flexscale, bpm*tonemulti*flexscale))

}


func flexMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){

				case 0: flexscale+=0.1

				}
			case 'h':
				switch(cursorpos){
					case 0:
						flexscale-=0.1


				}

			case 'L':
				switch(cursorpos){

				}
			case 'H':
				switch(cursorpos){


				}
	}
}
