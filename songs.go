package main

//import "fmt"

var songs []Song

var songsBoost = 1.0

type Song struct{
	clip SoundClip
	name string
	playing bool
	looping bool
	queued bool
}

func (s *Song) queue(){
	s.playing=false
	s.queued=true
}


func (s *Song) play(){
	s.clip.restartSound()
	s.playing=true
	s.queued=false
}

func playQueuedSongs(){

	for i := 0; i < len(songs); i++ {
		if(songs[i].queued){
			songs[i].play()
		}
	}

}



func Singing() float64 {


	var sum float64

	for i := 0; i < len(songs); i++ {
		if(songs[i].playing){
			songs[i].clip.speed = bpm/200.0
			sum+=Clip(songs[i].clip.playSound(), songs[i].clip.volume)


			if((int)(songs[i].clip.age)>len(songs[i].clip.wave)){
				songs[i].clip.restartSound()
				songs[i].playing=false

			}
		}

	}

	return sum*songsBoost



}
