package main

//import "fmt"

var samples []Sample

var samplesBoost = 1.0

type Sample struct{
	clip SoundClip
	name string
	playing bool
	interrupting bool
}


func (s *Sample) play(){
	s.clip.restartSound()
	s.interrupting=true
	s.playing=true
}

func (s *Sample) playGentle(){
	s.clip.restartSound()
	s.interrupting=false
	s.playing=true
}




func Sampling() float64 {


	var sum float64

	for i := 0; i < len(samples); i++ {
		if(samples[i].playing){
			sum+=Clip(samples[i].clip.playSound(), samples[i].clip.volume)


			if(samples[i].clip.age>len(samples[i].clip.wave)){
				samples[i].clip.restartSound()
				samples[i].playing=false

			}
		}

	}

	return sum*samplesBoost



}
