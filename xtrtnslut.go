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

var bpm float64 = 100
var tonemulti float64 = 4


var pausedrums bool = false
var toneall bool = false

var frozentones bool = false
var frozentonemulti float64 = 1


var drumsVolume float64 = 1
var songsVolume float64 = 1
var samplesVolume float64 = 1


func limitParameters(){

	if(bpm<40){bpm=40}
	if(bpm>4000){bpm=4000}

	if(tonemulti<1){tonemulti=1}
	if(tonemulti>1024){tonemulti=1024}

	if(editphase<0){editphase=0}
	if(editphase>len(phases)-1){editphase=len(phases)-1}

	if(editdrum<0){editdrum=0}
	if(editdrum>len(tones)-1){editdrum=len(tones)-1}

	if(editsample<0){editsample=0}
	if(editsample>len(samples)-1){editsample=len(samples)-1}

	if(editsong<0){editsong=0}
	if(editsong>len(songs)-1){editsong=len(songs)-1}


	if(feedback<0){feedback=0}
	if(feedback>1){feedback=1}
	if(reverbLength<0.04){reverbLength=0.04}

	if(reverbLength>4.9){reverbLength=4.9}

	for drumid := range tones{
		if(tones[drumid].multi<(1.0/32.0)){tones[drumid].multi=(1.0/32.0)}
		if(tones[drumid].multi>16){tones[drumid].multi=16}

		if(tones[drumid].depend< -1){tones[drumid].depend=-1}
		if(tones[drumid].depend>len(phases)-1){tones[drumid].depend=len(phases)-1}


		if(tones[drumid].clip.volume< 0.05){tones[drumid].clip.volume=0.05}
		if(tones[drumid].clip.volume> 3){tones[drumid].clip.volume=3}

		if(tones[drumid].clip.speed< 0.04){tones[drumid].clip.speed=0.04}

		if(tones[drumid].clip.speed> 5){tones[drumid].clip.speed=5}

	}

	for sampleid := range samples{

		if(samples[sampleid].clip.volume< 0.05){samples[sampleid].clip.volume=0.05}
		if(samples[sampleid].clip.volume> 3){samples[sampleid].clip.volume=3}

		if(samples[sampleid].clip.speed< 0.04){samples[sampleid].clip.speed=0.04}
		if(samples[sampleid].clip.speed> 5){samples[sampleid].clip.speed=5}

		if(samples[sampleid].loopfreq<0){samples[sampleid].loopfreq=0}
		if(samples[sampleid].loopfreq>4){samples[sampleid].loopfreq=4}

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


    ///--------------------


    //load drums:
	files, err := ioutil.ReadDir("./sounds/drums")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if(!file.IsDir()){

			f := strings.Split(file.Name(), ".")
			if len(f) > 0 {

				if(f[1]=="wav"){
					tones = append(tones, Tone{createSound("sounds/drums/"+file.Name()),false,false,1,f[0],-1,-1,0})
				}

			}

						//
		}
	}


	// song misc:
	songfiles, err := ioutil.ReadDir("./sounds/songs/")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range songfiles {
		if(!file.IsDir()){

			f := strings.Split(file.Name(), ".")
			if len(f) > 0 {

				if(f[1]=="wav"){
					songs = append(songs, Song{createSound("sounds/songs/"+file.Name()),f[0],false,false,false})
				}

			}

						//
		}
	}


	// load misc:
	miscfiles, err := ioutil.ReadDir("./sounds/misc/")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range miscfiles {
		if(!file.IsDir()){

			f := strings.Split(file.Name(), ".")
			if len(f) > 0 {

				if(f[1]=="wav"){
					samples = append(samples, Sample{createSound("sounds/misc/"+file.Name()),f[0],false,false,false,1,[4]bool{true,false,false,false}})
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

func writeState() string {
    result := ""

    if toneall {
        result += "[ALLTONES] "
    }

    if pausedrums {
        result += "[DRUMSPAUSED] "
    }


    if frozentones {
        result += "[FROZEN] "
    }

    if result == "" {
	result = "Running normally."
    }

    return result
}
