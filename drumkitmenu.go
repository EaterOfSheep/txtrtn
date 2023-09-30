package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var drumKitEdit = 0


func drumKitMenuDisplay(s *tcell.Screen, style tcell.Style){


	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Editing %s (%d/3)",drumName(drumKitEdit),drumKitEdit))
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Current: %s (%d/%d)",currentDrumName(drumKitEdit,0), currentDrumNumber(drumKitEdit,0), drumArraySize(drumKitEdit)))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Next: %s (%d/%d)",currentDrumName(drumKitEdit,1), currentDrumNumber(drumKitEdit,1), drumArraySize(drumKitEdit)))

/*
	drawText(*s, 5, 3, 50, 3, style, fmt.Sprintf("Current Kick: %d: %s",currentKick,tones[kickIds[currentKick]].name))

	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Current Clap: %d: %s",currentClap,tones[clapIds[currentClap]].name))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Current Hat: %d: %s",currentHat,tones[hatIds[currentHat]].name))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Current Snare: %d: %s",currentSnare,tones[snareIds[currentSnare]].name))
	*/

/*
	drawText(*s, 5, 4, 50, 4, style, fmt.Sprintf("Enabled: %v",tones[editdrum].playing))
	drawText(*s, 5, 5, 50, 5, style, fmt.Sprintf("Toning: %v",tones[editdrum].toning))
	drawText(*s, 5, 6, 50, 6, style, fmt.Sprintf("Tempo Multi: %.3f",tones[editdrum].multi))
	drawText(*s, 5, 7, 50, 7, style, fmt.Sprintf("Phase Dependence: %d",tones[editdrum].depend))
	drawText(*s, 5, 8, 50, 8, style, fmt.Sprintf("Phase Avoidance: %d",tones[editdrum].avoid))
	*/

}




func currentDrumNumber(d int, pos int) int{

	switch(pos){

		case 0:
			switch(d){
				case 0: return currentKick
				case 1: return currentSnare
				case 2: return currentClap
				case 3: return currentHat
			}

		case 1:
			switch(d){
				case 0: return nextKick
				case 1: return nextSnare
				case 2: return nextClap
				case 3: return nextHat
			}
	}






	return -1

}

func drumArraySize(d int) int{

		switch(d){
		case 0: return len(kickIds)-1
		case 1: return len(snareIds)-1
		case 2: return len(clapIds)-1
		case 3: return len(hatIds)-1

	}

	return 0

}

func currentDrumToneNumber(d int, pos int) int{

	switch(d){
		case 0: return kickIds[currentDrumNumber(d,pos)]
		case 1: return snareIds[currentDrumNumber(d,pos)]
		case 2: return clapIds[currentDrumNumber(d,pos)]
		case 3: return hatIds[currentDrumNumber(d,pos)]

	}

	return -1

}

func currentDrumName(d int, pos int) string{

	if(d<0){ return "Undrum" }

	return tones[currentDrumToneNumber(d,pos)].name

}

func drumName(d int) string{

	switch(d){
		case 0: return "Kick"
		case 1: return "Snare"
		case 2: return "Clap"
		case 3: return "Hat"

	}

	return "Fakedrum"

}


func editDrum(drum int){


	screenpos=0
	cursorpos=0


}

func setDrumKitDrum(d int, pos int, id int){

	switch(pos){

	case 0:
		switch(d){
			case 0: currentKick=id
			case 1: currentSnare=id
			case 2: currentClap=id
			case 3: currentHat=id
		}

	case 1:
		switch(d){
			case 0: nextKick=id
			case 1: nextSnare=id
			case 2: nextClap=id
			case 3: nextHat=id
		}
}


}


func drumKitSwitch(d int, pos int, change int){

	var output = currentDrumNumber(d,pos)

	output+=change

	if(output<0){output=drumArraySize(d)}

	if(output>drumArraySize(d)){output=0}

	setDrumKitDrum(d,pos,output)



}


func drumKitMenuControl(r rune){

	switch(r){





		case 'l':
			switch(cursorpos){
				case 0: if(drumKitEdit<3){drumKitEdit++}else{drumKitEdit=0}
				case 1: drumKitSwitch(drumKitEdit,0,+1)
				case 2: drumKitSwitch(drumKitEdit,1,+1)




			}
		case 'h':
			switch(cursorpos){
				case 0: if(drumKitEdit>0){drumKitEdit--}else{drumKitEdit=3}
				case 1: drumKitSwitch(drumKitEdit,0,-1)
				case 2: drumKitSwitch(drumKitEdit,1,-1)

			}

		case 'L':
			switch(cursorpos){

			}
		case 'H':
			switch(cursorpos){

			}

	}
}
