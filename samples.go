package main

//import "fmt"

var samples []Sample

var samplesBoost = 1.0

type Sample struct{
	clip SoundClip
	name string
	playing bool
	interrupting bool
	looping bool
	loopfreq int
}


func (s *Sample) play(){
	s.clip.restartSound()
	s.playing=true
}


func loopSamples(x int){

		for i := 0; i < len(samples); i++ {
			if(samples[i].looping && samples[i].loopfreq==x){
				samples[i].play()
			}
		}

}



func Sampling() float64 {


	var sum float64

	for i := 0; i < len(samples); i++ {
		if(samples[i].playing){
			sum+=Clip(samples[i].clip.playSound(), samples[i].clip.volume)


			if((int)(samples[i].clip.age)>len(samples[i].clip.wave)){
				samples[i].clip.restartSound()
				samples[i].playing=false

			}
		}

	}

	return sum*samplesBoost



}
