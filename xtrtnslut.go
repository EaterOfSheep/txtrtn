package main

import (
	"log"
	"io/ioutil"
	"strings"

	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

var superbar = 0
var bar = 0
var beat = 0
var beatstep = 0
var beatsubstep = 0
var beatage = 0

var bpm float64 = 180
var tonemulti float64 = 4


func limitParameters(){

	if(bpm<40){bpm=40}
	if(bpm>4000){bpm=4000}

	if(tonemulti<1){tonemulti=1}
	if(tonemulti>1024){tonemulti=1024}

	if(editphase<0){editphase=0}
	if(editphase>len(phases)-1){editphase=len(phases)-1}

	if(editdrum<0){editdrum=0}
	if(editdrum>len(tones)-1){editdrum=len(tones)-1}

	for drumid := range tones{
		if(tones[drumid].multi<0.1){tones[drumid].multi=0.1}
		if(tones[drumid].multi>8){tones[drumid].multi=8}

		if(tones[drumid].depend< -1){tones[drumid].depend=-1}
		if(tones[drumid].depend>len(phases)-1){tones[drumid].depend=len(phases)-1}


		if(tones[drumid].clip.volume< 0.55){tones[drumid].clip.volume=0.5}

		if(tones[drumid].clip.volume> 2.95){tones[drumid].clip.volume=3}

	}


	for phaseid := range phases{
		if(phases[phaseid].level<0){phases[phaseid].level=0}
		if(phases[phaseid].level>4){phases[phaseid].level=4}

		if(phases[phaseid].tonemulti<1){phases[phaseid].tonemulti=1}
		if(phases[phaseid].tonemulti>8){phases[phaseid].tonemulti=8}
		if(phases[phaseid].basemulti<1){phases[phaseid].basemulti=1}
		if(phases[phaseid].basemulti>8){phases[phaseid].basemulti=8}

		if(phases[phaseid].depend< -1){phases[phaseid].depend=-1}
		if(phases[phaseid].depend>phaseid-1){phases[phaseid].depend=phaseid-1}
		if(phases[phaseid].avoid< -1){phases[phaseid].avoid=-1}
		if(phases[phaseid].avoid>phaseid-1){phases[phaseid].avoid=phaseid-1}
	}

}

func main() {


	files, err := ioutil.ReadDir("./sounds")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if(!file.IsDir()){

			f := strings.Split(file.Name(), ".")
			if len(f) > 0 {

				if(f[1]=="wav"){
					tones = append(tones, Tone{createSound("sounds/"+file.Name()),false,false,1,f[0],-1,-1})
				}

			}

						//
		}
	}


	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(XTRTN())
	defer speaker.Close()

	for drumid := range tones{
		tones[drumid].clip.age = 44100*3 //3 seconds wait
	}

	for phaseid := range phases{
		phases[phaseid].avoid=-1
		phases[phaseid].depend=-1
		phases[phaseid].alive=true
	}

	showterminal()

}
