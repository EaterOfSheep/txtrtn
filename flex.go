package main

import (
	"fmt"
	"math"
)
var flexing = false
var frozentoneflex bool = false
var flexratio float64 = 1
var flexscale float64 = 1
var flextype int = 0

func writeFlexing() string{

	if !flexing {return "No flex"}


	return "Flex: "+flexTypeName()+" | "+fmt.Sprintf("Ratio: x%v",flexratio)

}

func flexTypeName() string{

	switch(flextype){

		case 0: return "Sawtooth"
		case 1: return "Sine"
	}

	return "Unknown"

}


func updateFlexRatio() {

	flexratio = 1+(flexProcess(deepCompletionRatio())*flexscale)

}

func flexProcess(x float64) float64{

		if(!flexing) {return 0}

		if frozentones && !frozentoneflex {return 0}

		switch(flextype){

		case 0: return x
		case 1: return math.Sin(x*math.Pi*2)+1

		}

		return 0

}

func deepCompletionRatio() float64{

	return (completionRatio(beatstep)*0.25*0.25*0.25+completionRatio(beat)*0.25*0.25+completionRatio(bar)*0.25+completionRatio(superbar))

	//To do: make this a recursive function with levels

}


func completionRatio(x int) float64 {
	return ((float64)(x)/4)
}
