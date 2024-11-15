package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func mainMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Base Tempo: %vbpm",bpm))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Tone Multiplier: x%v (%vbpm)",tonemulti, bpm*tonemulti))

	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Pause drums (z): %v",pausedrums))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Tone Everything (t): %v",toneall))
	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Frozen Tones (f): %v",frozentones))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Frozen Tone Multiplier: x%v (%vbpm)",frozentonemulti, bpm*tonemulti*frozentonemulti))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Frozen Tone Flex: %v",frozentoneflex))



	drawText(*s, 5, 10, 50, 10, style, fmt.Sprintf("Boost Tones (x): %v",boosttones))
	drawText(*s, 5, 11, 50, 11, style, fmt.Sprintf("Boost Tone Multiplier: x%v (%vbpm)",boosttonemulti, bpm*tonemulti*boosttonemulti))

	drawText(*s, 5, 12, 50, 12, style, fmt.Sprintf("Boost Tone Slew Up: %.2f",boosttoneupslew))

	drawText(*s, 5, 13, 50, 13, style, fmt.Sprintf("Boost Tone Slew Down: %.2f",boosttonedownslew))


}


func mainMenuControl(r rune){

	switch(r){

		case 'l':
			switch(cursorpos){
				case 0:
					bpm+=5
				case 1:
					tonemulti+=1

				case 2: pausedrums=!pausedrums

				case 3: toneall=!toneall

				case 4: frozentones=!frozentones
				case 5: frozentonemulti+=1
				case 6: frozentoneflex=!frozentoneflex
				case 7: boosttones=!boosttones
				case 8: boosttonemulti+=1
				case 9: boosttoneupslew+=0.02
				case 10: boosttonedownslew+=0.1


					/*
				case 2:

				case 3:
					editphase++
				case 4:
					phases[editphase].level++
				case 5:
					phases[editphase].location[0]=!phases[editphase].location[0]
				case 6:
					phases[editphase].location[1]=!phases[editphase].location[1]
				case 7:
					phases[editphase].location[2]=!phases[editphase].location[2]
				case 8:
					phases[editphase].location[3]=!phases[editphase].location[3]
				case 9:
					phases[editphase].depend++
				case 10:
					phases[editphase].avoid++


				case 11:
					phases[editphase].tonemulti++
					*/



				}
			case 'h':
				switch(cursorpos){
					case 0:
						bpm-=5
					case 1:
						tonemulti-=1

					case 5:
						frozentonemulti-=1

					case 8: boosttonemulti-=1
					case 9: boosttoneupslew-=0.02
					case 10: boosttonedownslew-=0.1

						/*
					case 3:
						editphase--

					case 4:
						phases[editphase].level--
					case 9:
						phases[editphase].depend--
					case 10:
						phases[editphase].avoid--
					case 11:
						phases[editphase].tonemulti--
						*/

				}

			case 'L':
				switch(cursorpos){
					case 0:
						bpm+=100
					case 1:
						tonemulti*=2

					case 5:
						frozentonemulti*=2

					case 8:
						boosttonemulti*=2

				}
			case 'H':
				switch(cursorpos){
					case 0:
						bpm-=100
					case 1:
						tonemulti/=2

					case 5:
						frozentonemulti/=2

					case 8:
						boosttonemulti/=2

				}
	}
}


func helpDisplay(s *tcell.Screen, style tcell.Style){


	drawText(*s, 5, 3, 100, 3, style, fmt.Sprintf("This is Terminal Extratone (TXTRTN)."))
	drawText(*s, 5, 4, 100, 4, style, fmt.Sprintf("Made by Eater of Sheep in 2022."))

	drawText(*s, 5, 6, 100, 6, style, fmt.Sprintf("TXTRTN is a text-based, vimcentric DAW for performing live Extratone."))

	drawText(*s, 5, 8, 100, 8, style, fmt.Sprintf("Use j/k to move the cursor."))
	drawText(*s, 5, 9, 100, 9, style, fmt.Sprintf("Use h/l to interact with menu items."))
	drawText(*s, 5, 10, 100, 10, style, fmt.Sprintf("Press 0 - 9 to switch between menus."))
	drawText(*s, 5, 11, 100, 11, style, fmt.Sprintf("You can return to this guide by pressing w."))

	drawText(*s, 5, 13, 100, 13, style, fmt.Sprintf("Press q to exit."))


}
