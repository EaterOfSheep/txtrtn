package main

//import "fmt"
import (
	"math/rand"
	"time"
)

var tones []Tone
var kickIds []int
var snareIds []int
var clapIds []int
var hatIds []int

var currentKick = 0
var currentSnare = 0
var currentHat = 0
var currentClap = 0

var nextKick = 1
var nextSnare = 1
var nextHat = 1
var nextClap = 1

var kickPlaying = false
var snarePlaying = false
var hatPlaying = false
var clapPlaying = false

var drumKitToning = [4]bool{false, false, false, false}
var drumKitMulti = [4]float64{1, 1, 1, 1}


type Tone struct{
	clip SoundClip
	playing bool
	toning bool
	multi float64
	name string
	depend int
	avoid int
	trueage int
}

func clearTones(){

	for i := range tones{

		//tones[i].playing=false
		//tones[i].toning=false
		//tones[i].multi=1
		tones[i].depend=-1
		tones[i].avoid=-1

	}

}

func syncDrums(){
		for i := 0; i < len(tones); i++ {
			if(tones[i].playing&&!((int)(tones[i].trueage)<len(tones[i].clip.wave))){
				tones[i].trueage=44100*3
				tones[i].clip.age=44100*3
			}
		}
}

func canTonePlay(t Tone) bool{

	if(!t.playing){return false}

	if(t.avoid!=-1){
		if(phases[t.avoid].on){return false}
	}

	if(t.depend==-1){return true}




	return phases[t.depend].on

}

func shouldToneBePlaying(id int) bool{

		if(kickIds[currentKick]==id && kickPlaying){
			return true
		}

		if(snareIds[currentSnare]==id && snarePlaying){
			return true
		}

		if(clapIds[currentClap]==id && clapPlaying){
			return true
		}

		if(hatIds[currentHat]==id && hatPlaying){
			return true
		}

		return false


}

func enforceDrumKitRules(id int){

		tones[id].playing=shouldToneBePlaying(id)

		for i:=0;i<4;i++ {

			if(currentDrumToneNumber(i,0)==id){
				tones[id].multi = drumKitMulti[i]
				tones[id].toning = drumKitToning[i]
			}
		}


}

func sbarAutomation(){

	rand.Seed(time.Now().UnixNano())

	if(autoBoosting){
		boosttones=(rand.Intn(100)<boostingChance)
	}
	if(autoFreezing){
		frozentones=(rand.Intn(100)<freezingChance)
	}
	if(autoFlexing){
		flexing=(rand.Intn(100)<flexingChance)
	}

}




func Drumming() float64{

	beatage++

	if(beatage>SecondsToFrame((60/16)/bpm)){//quarter step

		updateToneBoost()
		loopSamples(4,beatsubstep)
		beatage=0
		beatsubstep++
		updateFlexRatio()
		//redrawTerminal()
	}

	if(beatsubstep>3){
		loopSamples(3,beatstep)
		beatsubstep=0
		beatstep++
	}

	if(beatstep>3){
		loopSamples(2,beat)
		beatstep=0
		beat++
	}

	if(beat>3){
		loopSamples(1,bar)
		beat=0
		bar++
	}

	if(bar>3){

		sbarAutomation()

		loopSamples(0,superbar)
		bar=0
		bpm+=bpmAutoClimbSbar
		superbar++
		if(autoregen || autopush){
			go autoGenTimer()
		}
		syncDrums() //sync on superbar for now
		playQueuedSongs() //play queued songs on superbar
	}

	if(superbar>3){
		superbar=0

	}

	var sum float64

	for i := 0; i < len(tones); i++ {

		enforceDrumKitRules(i)



		sum+=Clip(tones[i].clip.playSound(), tones[i].clip.volume)
		tones[i].trueage++
		truebpm := bpm*tones[i].multi
		if(tones[i].toning || toneall){
			truebpm*=tonemulti*getPhaseToneMulti()
			truebpm*=flexratio
			truebpm*=trueBoostToneMulti()
		}else{
			truebpm*=getPhaseBaseMulti()
		}

		if(canTonePlay(tones[i])&& (int)(tones[i].trueage)>SecondsToFrame(60/truebpm)){

			var samplingnow = false

			for j := 0; j < len(samples); j++ {

				if(samples[j].playing && samples[j].interrupting){samplingnow=true}

			}

			if(!samplingnow && !pausedrums){tones[i].trueage = 0; tones[i].clip.restartSound()}
		}

	}

	return sum

}

func writeDrums() string {
    result := ""

    if kickPlaying {
	    if(drumKitToning[0]){
		result+="[K-TONE] "
	    }else{
		result += "[KICK] "
	    }
    }

    if snarePlaying {
	    if(drumKitToning[1]){
		result+="[S-TONE] "
	    }else{
		result += "[SNARE] "
	    }
    }

    if clapPlaying {
	    if(drumKitToning[2]){
		result+="[C-TONE] "
	    }else{
		result += "[CLAP] "
	    }
    }

    if hatPlaying {
	    if(drumKitToning[3]){
		result+="[H-TONE] "
	    }else{
		result += "[HAT] "
	    }
    }



    if result == "" {
	result = "No active drums."
    }

    return result
}
