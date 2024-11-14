package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)



func automationMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Tempo Climb (On AutoGen): %vbpm",bpmAutoClimb))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Tempo Climb On New Sbar: %vbpm",bpmAutoClimbSbar))


	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Autofreezing: %v",autoFreezing))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Autofreeze Rate: %v%%",freezingChance))
	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Autoboosting: %v",autoBoosting))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Autoboost Rate: %v%%",boostingChance))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Autoflexing: %v",autoFlexing))
	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Autoflex Rate: %v%%",flexingChance))


	drawText(*s, 5, 11, 50, 11, style, fmt.Sprintf("Auto Tone Multi: %v%%",autoToneMultiChance))

}


func automationMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){
				case 0:
					if(bpmAutoClimb<20){bpmAutoClimb++}
				case 1:
					if(bpmAutoClimbSbar<20){bpmAutoClimbSbar++}

				case 2: autoFreezing = !autoFreezing

				case 3:
					if(freezingChance<100){freezingChance++}

				case 4: autoBoosting = !autoBoosting

				case 5:
					if(boostingChance<100){boostingChance++}



				case 6: autoFlexing = !autoFlexing

				case 7:
					if(flexingChance<100){flexingChance++}
				case 8:
					if(autoToneMultiChance<100){autoToneMultiChance++}



			}
		case 'h':
			switch(cursorpos){

				case 0:
					if(bpmAutoClimb>0){bpmAutoClimb--}
				case 1:
					if(bpmAutoClimbSbar>0){bpmAutoClimbSbar--}

				case 3:
					if(freezingChance>0){freezingChance--}
				case 5:
					if(boostingChance>0){boostingChance--}
				case 7:
					if(flexingChance>0){flexingChance--}
				case 8:
					if(autoToneMultiChance>0){autoToneMultiChance--}



			}


	}
}
