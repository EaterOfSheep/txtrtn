package main


const reverbBufferSize int = 14000
var reverbBuffer [reverbBufferSize]float64
var feedback float64 = 0.5
var reverbPosition int = 0



func reverb(input float64) float64 {
    output := input + reverbBuffer[reverbPosition]*feedback
    reverbBuffer[reverbPosition] = output
    reverbPosition = (reverbPosition + 1) % reverbBufferSize
    return output
}
