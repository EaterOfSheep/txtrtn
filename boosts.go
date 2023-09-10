package main

var boosttones bool = false
var boosttonemulti float64 = 4

var boosttoneprogress float64 = 0
var boosttoneupslew float64 = 0
var boosttonedownslew float64 = 0

func updateToneBoost(){

	var goal float64 = -1
	var slew = boosttonedownslew

	if(boosttones){
		goal = 1
		slew = boosttoneupslew
	}

	if(slew!=0){
		goal/=64
		goal/=slew
	}

	boosttoneprogress+=goal

	if(boosttoneprogress>1){boosttoneprogress=1}

	if(boosttoneprogress<0){boosttoneprogress=0}


}

func trueBoostToneMulti() float64{

	return 1+(boosttonemulti-1)*boosttoneprogress

}
