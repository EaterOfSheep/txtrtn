package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)



func automationMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Tempo Climb On Pattern AutoGen: %vbpm",bpmAutoClimb))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Tempo Climb On New Sbar: %vbpm",bpmAutoClimbSbar))

}


func automationMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){
				case 0:
					if(bpmAutoClimb<20){bpmAutoClimb++}
				case 1:
					if(bpmAutoClimbSbar<20){bpmAutoClimbSbar++}

			}
		case 'h':
			switch(cursorpos){

				case 0:
					if(bpmAutoClimb>0){bpmAutoClimb--}
				case 1:
					if(bpmAutoClimbSbar>0){bpmAutoClimbSbar--}

			}


	}
}
