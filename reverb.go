package main



var feedback float64 = 0.3
var reverbPosition int = 0
var reverbLength float64 = 0.2 //max 5 as defined by reverbBufferSize

const reverbBufferSize int = 44100*5
var reverbBuffer [reverbBufferSize]float64



func reverb(input float64) float64 {
    output := input + reverbBuffer[reverbPosition]*feedback
    reverbBuffer[reverbPosition] = output
    reverbPosition = (reverbPosition + 1) % SecondsToFrame(reverbLength)
    return output
}
