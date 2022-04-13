package main

import (
	"math/rand"
	"time"
)

var genselect int
var autoregen bool
var autoregencount=3
var autoregencountmax=3
var autoregens = make(chan bool)

var gennumber=0

var tonebursts int
var basebursts int


func clearPatterns(){

	clearPhases()
	clearTones()

}

func addPhase(level int, location [4]bool, depend int, avoid int) int{
	genselect++
	phases[genselect].location = location
	phases[genselect].level = level
	phases[genselect].depend = depend
	phases[genselect].avoid = avoid
	return genselect
}

func autoGenTimer(){

	if(autoregen){
		if(autoregencount==0){
			autoregencount=autoregencountmax
			autoregens<-true
			//autoregen=false
		}else{
			autoregencount--
		}
	}

}

func addBurst(basemulti float64, tonemulti float64, levels []int, depend int, avoid int) int{

	for i:=range levels{
		genselect++
		phases[genselect].level = levels[i]
		if(i!=0){phases[genselect].depend=genselect-1}else{
			phases[genselect].depend = depend
			phases[genselect].avoid = avoid
		}
		phases[genselect].location=randomPhaseLocs(1,3)
	}

	phases[genselect].basemulti = basemulti
	phases[genselect].tonemulti = tonemulti
	return genselect
}

func generatePatterns(){
	gennumber++

	rand.Seed(time.Now().UnixNano())
	genselect = -1
	clearPatterns()

	//---
	//levels: 0 sbar, 1 bar, 2 beat, 3 step, 4 sstep

	lastbar:=addPhase(1,[4]bool{false,false,false,true}, -1,-1)

	for i := range tones{
		tones[i].avoid=addBurst(0,0, []int{0,2} ,lastbar,-1)
		//addPhase(2,randomPhaseLocs(1,3),lastbar,-1)
	}

	//superbars1:=addBurst(2,0, []int{0} ,lastbar,-1)
	//superbars2:=addBurst(2,0, []int{0} ,lastbar,-1) //random possibly groups of superbars

	//addBurst(2,2, []int{2,3} ,lastbar,superbars1)
	//addBurst(2,2, []int{2,3} ,lastbar,superbars2)

	for i:=0; i<tonebursts; i++{
		addBurst(0,2, []int{2,3} ,-1,-1)
	}

	for i:=0; i<basebursts; i++{
		addBurst(2,0, []int{2,3} ,-1,-1)
	}


}

func randomPhaseLocs(min int, max int) [4]bool{ //maximum and minumum numbers of active locations

	activeLocs:=rand.Intn(1+max-min)+min

	if(activeLocs<1){return [4]bool{false,false,false,false}}
	if(activeLocs>3){return [4]bool{true,true,true,true}}

	var selectedLocs []int
	output := [4]bool{false,false,false,false}
	for i := 0; i<activeLocs; i++{
		solved:=false
		selectedLocs=append(selectedLocs,-1)
		for !solved{
			selectedLocs[i]=rand.Intn(4) //0,1,2,3
			solved=true

			for j := 0; j<i; j++{
				if(selectedLocs[j]==selectedLocs[i]){
					solved=false
				}
			}

		}
		output[selectedLocs[i]]=true
	}

	return output

}
