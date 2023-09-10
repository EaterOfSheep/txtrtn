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

	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Playback Speed: x%.2f",samples[editsample].clip.speed))


	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Looping: %t",samples[editsample].looping))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Loop Level: %s",sampleLoopFreqName(samples[editsample])))
	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Loop Location 0: %t",samples[editsample].location[0]))
	drawText(*s, 5, 11, 50, 11, style, fmt.Sprintf("Loop Location 1: %t",samples[editsample].location[1]))
	drawText(*s, 5, 12, 50, 12, style, fmt.Sprintf("Loop Location 2: %t",samples[editsample].location[2]))
	drawText(*s, 5, 13, 50, 13, style, fmt.Sprintf("Loop Location 3: %t",samples[editsample].location[3]))


	drawText(*s, 5, 15, 50,15, style, fmt.Sprintf("Length: %.2fs [%.2f%%]",(float64)(len(samples[editsample].clip.wave))/44100.0, intPercentage((int)(samples[editsample].clip.age),len(samples[editsample].clip.wave))))



}


func sampleLoopFreqName(s Sample) string{

	switch(s.loopfreq){
		case 0: return "Superbar"
		case 1: return "Bar"
		case 2: return "Beat"
		case 3: return "Step"
		case 4: return "Substep"
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
					samples[editsample].clip.speed+=0.02


				case 5:
					samples[editsample].looping = !samples[editsample].looping
				case 6:
					samples[editsample].loopfreq++

				case 7:
					samples[editsample].location[0]=!samples[editsample].location[0]
				case 8:
					samples[editsample].location[1]=!samples[editsample].location[1]
				case 9:
					samples[editsample].location[2]=!samples[editsample].location[2]
				case 10:
					samples[editsample].location[3]=!samples[editsample].location[3]


			}
		case 'h':
			switch(cursorpos){
				case 0:
					editsample--
				case 3:
					samples[editsample].clip.speed-=0.02
				case 6:
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
