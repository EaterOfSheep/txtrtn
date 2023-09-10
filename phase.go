package main

import "fmt"

var phases [64]Phase

type Phase struct{
	tonemulti float64
	basemulti float64
	level int //0=superbar 1=bar, 2=beat, 3=step, 4=substep
	location [4]bool //0-3
	on bool
	alive bool
	depend int
	avoid int
}

func clearPhases(){
	for i := range phases{
		phases[i].tonemulti=1
		phases[i].basemulti=1
		phases[i].level=0
		phases[i].location=[4]bool{false,false,false,false}
		phases[i].on=false
		phases[i].alive=true
		phases[i].depend=-1
		phases[i].avoid=-1
	}

}

func getPhaseDependence(p Phase) []int{

	var list []int
	current := p.depend

	for current!=-1 {
		list=append(list,current)
		current=phases[current].depend
	}

	return list
}

func getPhaseAvoidance(p Phase) []int{

	var list []int
	if(p.avoid!=-1){
		list=append(list,p.avoid)
	}

	for _, i := range getPhaseDependence(p) {
		if(phases[i].avoid!=-1){
			list=append(list,phases[i].avoid)
		}
	}

	return list
}

func getPhaseLevelName(level int) string{

	switch(level){
		case 0:
			return "superbar"
		case 1:
			return "bar"
		case 2:
			return "beat"
		case 3:
			return "step"
		case 4:
			return "substep"
	}
	return "error"


}

func getPhaseBaseMulti() float64{

	var product float64 = 1

	for _, phase := range(phases){
		if(phase.on){
			product*=phase.basemulti
		}
	}

	return product

}

func getPhaseToneMulti() float64{

	if frozentones {return frozentonemulti}


	var product float64 = 1

	for _, phase := range(phases){
		if(phase.on){
			product*=phase.tonemulti
		}
	}

	return product

}

func Phasings() {

	for phaseid, phase := range(phases){

		var on bool

		switch(phase.level){
			case 0:
				on = phase.location[superbar]
			case 1:
				on = phase.location[bar]
			case 2:
				on = phase.location[beat]
			case 3:
				on = phase.location[beatstep]
			case 4:
				on = phase.location[beatsubstep]
		}

		if(phase.avoid!=-1){
			if(phases[phase.avoid].on){
				on=false
			}
		}
		if(phase.depend!=-1){
			if(!phases[phase.depend].on){
				on=false
			}
		}

		if(!phase.alive){
			on=false
		}

		phases[phaseid].on = on
	}
}

func writePhases() string{
	text := "Active Phases: "
	for phaseid := range phases {
		if(phases[phaseid].on){
			text=fmt.Sprintf("%s %d",text, phaseid)
		}
	}
	return text

}
