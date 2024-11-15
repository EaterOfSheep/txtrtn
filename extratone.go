package main

import (
	"os"
	"io"

	"github.com/youpy/go-wav"
	"github.com/faiface/beep"
)



type SoundClip struct{
	wave []float64
	volume float64
	age float64
	speed float64
}

func createSound(wavpath string) SoundClip{

	return SoundClip{openWav(wavpath),1,0,1}

}

func openWav(wavpath string) []float64{

	file, _ := os.Open(wavpath)
	reader := wav.NewReader(file)

  	defer file.Close()

	var outputwav []float64
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			outputwav=append(outputwav,reader.FloatValue(sample, 0))
		}
	}

	return outputwav


}


func SecondsToFrame(sec float64) int{
	return int(44100*sec)
}

func XTRTN() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {

		Phasings()

		for i := range samples {

			//limitParameters()
			displacement := reverb(Drumming()*drumsVolume+Sampling()*samplesVolume+Singing()*songsVolume)


		//	displacement := Clip(playSound(kickwav, t) + playSound(snarewav, t),distortion)


			samples[i][0] = displacement
			samples[i][1] = displacement
		}
		return len(samples), true
	})
}

func Clip(x float64, scale float64) float64{
	x*=scale
	if(x > 1){
		x=1
	}

	if(x < -1){
		x=-1
	}
	return x
}

func (clip *SoundClip) restartSound(){
	clip.age=0
}

func (clip *SoundClip) playSound() float64{

			clip.age+=clip.speed

			//freq := 600*math.Exp(-(decayspeed*t))
			//return math.Sin(freq*t*math.Pi)
			if((int)(clip.age)<len(clip.wave)-1){
				return clip.wave[(int)(clip.age)]
			}else{
				return 0
			}

}
