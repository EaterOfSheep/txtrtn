package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var editsample int = 0

func sampleMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Sample %d: %s",editsample,samples[editsample].name))
	drawText(*s, 5, 4, 50, 4, style, "Play! (Space)")
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Background: %t",!samples[editsample].interrupting))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Looping: %t",samples[editsample].looping))

	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Looping: %s",sampleLoopFreqName(samples[editsample])))


	drawText(*s, 5, 9, 50,9, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(samples[editsample].clip.wave))/44100.0, intPercentage((int)(samples[editsample].clip.age),len(samples[editsample].clip.wave))))



}


func sampleLoopFreqName(s Sample) string{

	switch(s.loopfreq){
		case 0: return "4Superbar"
		case 1: return "Superbar"
		case 2: return "Bar"
		case 3: return "Beat"
		case 4: return "Step"
		case 5: return "Substep"
	}
	return "???"
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
				case 2:
					samples[editsample].interrupting = !samples[editsample].interrupting
				case 3:
					samples[editsample].looping = !samples[editsample].looping
				case 4:
					samples[editsample].loopfreq++


			}
		case 'h':
			switch(cursorpos){
				case 0:
					editsample--
				case 4:
					samples[editsample].loopfreq--


			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
