package main

//import "fmt"

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

var kickPlaying = true
var snarePlaying = true
var hatPlaying = true
var clapPlaying = true


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

		loopSamples(0,superbar)
		bar=0
		superbar++
		if(autoregen){
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

		tones[i].playing=shouldToneBePlaying(i)


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
