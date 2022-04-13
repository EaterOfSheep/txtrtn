package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var editphase int = 0

func phaseMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Editing Phase: %d",editphase))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Enabled: %t",phases[editphase].alive))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Level: %s",getPhaseLevelName(phases[editphase].level)))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Location 0: %t",phases[editphase].location[0]))
	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Location 1: %t",phases[editphase].location[1]))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Location 2: %t",phases[editphase].location[2]))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Location 3: %t",phases[editphase].location[3]))

	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Phase Depend: %d%v", phases[editphase].depend, getPhaseDependence(phases[editphase])))
	drawText(*s, 5, 11, 50, 11, style, fmt.Sprintf("Phase Avoid: %d%v",phases[editphase].avoid,getPhaseAvoidance(phases[editphase])))
	drawText(*s, 5, 12, 50, 12, style, fmt.Sprintf("Tonemulti: %v",phases[editphase].tonemulti))
	drawText(*s, 5, 13, 50, 13, style, fmt.Sprintf("Basemulti: %v",phases[editphase].basemulti))

}




func phaseMenuControl(r rune){

	switch(r){


		case 'K':
			editphase--
		case 'J':
			editphase++


		case 'l':
			switch(cursorpos){
				case 0:
					editphase++
				case 1:
					phases[editphase].alive = !phases[editphase].alive
				case 2:
					phases[editphase].level++
				case 3:
					phases[editphase].location[0]=!phases[editphase].location[0]
				case 4:
					phases[editphase].location[1]=!phases[editphase].location[1]
				case 5:
					phases[editphase].location[2]=!phases[editphase].location[2]
				case 6:
					phases[editphase].location[3]=!phases[editphase].location[3]
				case 7:
					if(phases[editphase].depend<editphase-1){
						phases[editphase].depend++
					}
				case 8:
					if(phases[editphase].avoid<editphase-1){
						phases[editphase].avoid++
					}


				case 9:
					phases[editphase].tonemulti++

				case 10:
					phases[editphase].basemulti++



				}
			case 'h':
				switch(cursorpos){
					case 0:
						editphase--
					case 2:
						phases[editphase].level--
					case 7:

						if(phases[editphase].depend>-1){
							phases[editphase].depend--
						}
					case 8:

						if(phases[editphase].avoid>-1){
							phases[editphase].avoid--
						}
					case 9:
						phases[editphase].tonemulti--
					case 10:
						phases[editphase].basemulti--

				}

			case 'L':
				switch(cursorpos){

				}
			case 'H':
				switch(cursorpos){

				}
	}
}
