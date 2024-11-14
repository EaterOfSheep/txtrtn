package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)



func genMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Clear Patterns"))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Generate #%d",gennumber))


	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Basebursts %d",basebursts))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Tonebursts %d",tonebursts))


	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Enable Auto-pattern-generate: %t",autoregen))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Autocount: %d/%d superbars",autoregencount,autoregencountmax))
	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Muting patterns: %t",generateAvoids))


	drawText(*s, 5, 11, 50, 11, style, fmt.Sprintf("Enable Auto-drum-push: %t",autopush))
	drawText(*s, 5, 12, 50, 12, style, fmt.Sprintf("Autocount: %d/%d superbars",autopushcount,autopushcountmax))
	drawText(*s, 5, 13, 50, 13, style, fmt.Sprintf("Drum Conservation: %d%%",drumConservation))


	drawText(*s, 5, 15, 50, 15, style, fmt.Sprintf("Tempo Climb: %vbpm/Sbar",bpmAutoClimb))


}


func genMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){
				case 0:
					clearPatterns()
				case 1:
					//autoregens<-true
					generatePatterns()
				case 2:
					if(basebursts<9){basebursts++}
				case 3:
					if(tonebursts<9){tonebursts++}

				case 5:
					autoregen=!autoregen
				case 6:
					if(autoregencount<31){autoregencount++}

				case 7:
					generateAvoids=!generateAvoids

				case 8:
					autopush=!autopush
				case 9:
					if(autopushcount<31){autopushcount++}
				case 10:
					if(drumConservation<100){drumConservation++}

				case 12:
					if(bpmAutoClimb<100){bpmAutoClimb++}



			}
		case 'h':
			switch(cursorpos){
				case 2:
					if(basebursts>0){basebursts--}
				case 3:
					if(tonebursts>0){tonebursts--}

				case 6:
					if(autoregencount>0){autoregencount--}

				case 9:
					if(autopushcount>0){autopushcount--}
				case 10:
					if(drumConservation>0){drumConservation--}
				case 12:
					if(bpmAutoClimb>0){bpmAutoClimb--}


			}

		case 'L':
			switch(cursorpos){

				case 6:
					if(autoregencountmax<31){autoregencountmax++}

				case 9:
					if(autopushcountmax<31){autopushcountmax++}



			}
		case 'H':
			switch(cursorpos){

				case 6:
					if(autoregencountmax>0){autoregencountmax--}
				case 9:
					if(autopushcountmax>0){autopushcountmax--}

			}
	}
}
