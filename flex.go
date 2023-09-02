package main

import (
	"fmt"
)

var flexratio float64 = 1
var flexscale float64 = 1

func writeFlexing() string{

	return fmt.Sprintf("Flex Ratio: x%v",flexratio)

}


func updateFlexRatio() {

	flexratio = 1+(deepCompletionRatio()*flexscale)


}

func deepCompletionRatio() float64{

	return (completionRatio(beatstep)*0.25*0.25)+completionRatio(beat)*0.25+(completionRatio(bar))

	//To do: make this a recursive function with levels

}


func completionRatio(x int) float64 {
	return ((float64)(x)/4)
}
