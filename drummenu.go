package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var editdrum int = 0

func drumMenuDisplay(s *tcell.Screen, style tcell.Style){

	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Editing Drum: %d: %s",editdrum,tones[editdrum].name))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Enabled: %v",tones[editdrum].playing))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Toning: %v",tones[editdrum].toning))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Tempo Multi: %v",tones[editdrum].multi))
	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Phase Dependence: %d",tones[editdrum].depend))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Phase Avoidance: %d",tones[editdrum].avoid))
	drawText(*s, 5, 9, 50, 9, style, fmt.Sprintf("Volume: %.2f",tones[editdrum].clip.volume))

}




func drumMenuControl(r rune){

	switch(r){


		case 'K':
			editdrum--
		case 'J':
			editdrum++


		case 'l':
			switch(cursorpos){
				case 0:
					editdrum++
				case 1:
					tones[editdrum].playing=!tones[editdrum].playing
				case 2:
					tones[editdrum].toning=!tones[editdrum].toning

				case 3:
					if(tones[editdrum].multi>0.95){
						tones[editdrum].multi++
					}else{
						tones[editdrum].multi+=0.1
					}
				case 4:
					if(tones[editdrum].depend<len(phases)-1){
						tones[editdrum].depend++
					}
				case 5:
					if(tones[editdrum].avoid<len(phases)-1){
						tones[editdrum].avoid++
					}
				case 6:
					tones[editdrum].clip.volume+=0.05



			}
		case 'h':
			switch(cursorpos){
				case 0:
					editdrum--

				case 3:
					if(tones[editdrum].multi>1){
						tones[editdrum].multi--
					}else{
						tones[editdrum].multi-=0.1
					}
				case 4:
					if(tones[editdrum].depend>-1){
						tones[editdrum].depend--
					}
				case 5:
					if(tones[editdrum].avoid>-1){
						tones[editdrum].avoid--
					}
				case 6:
					tones[editdrum].clip.volume-=0.05
			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
