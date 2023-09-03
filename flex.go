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
var flexFreq int = 1

func writeFlexing() string{

	if !flexing {return "No flex"}


	return "Flex: "+flexTypeName()+" | "+fmt.Sprintf("Ratio: x%v",flexratio)

}

func flexTypeName() string{

	switch(flextype){


		case 0: return "Sine"

		case 1: return "Sawtooth"

		case 2: return "Triangle"
	}

	return "Unknown"

}


func updateFlexRatio() {

	flexratio = 1+(flexProcess(variableCompletionRatio())*flexscale)

}

func flexProcess(x float64) float64{

		if(!flexing) {return 0}

		if frozentones && !frozentoneflex {return 0}

		switch(flextype){

		case 0: return (math.Sin(x*math.Pi*2)/2)+1
		case 1: return x
		case 2: if x < 0.5 {return 2*x} else {return 1-((x-0.5)*2)}

		}

		return 0

}

func variableCompletionRatio() float64{

	return math.Mod(deepCompletionRatio() * ((float64)(flexFreq)),1.0)

}

func deepCompletionRatio() float64{

	return (completionRatio(beatstep)*0.25*0.25*0.25+completionRatio(beat)*0.25*0.25+completionRatio(bar)*0.25+completionRatio(superbar))

	//To do: make this a recursive function with levels

}


func completionRatio(x int) float64 {
	return ((float64)(x)/4)
}
