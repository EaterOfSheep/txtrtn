package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var screenpos int = 0
var cursorpos int = 0

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drumAges() string{
	text := "Sounds: "
	for _, tone := range tones {
		text=fmt.Sprintf("%s %d/%d",text, tone.clip.age, len(tone.clip.wave))
	}
	return text

}

func drumNames() string{
	text := "Sounds: "
	for _, tone := range tones {
		text=fmt.Sprintf("%s %s",text, tone.name)
	}
	return text

}



func writeTimings() string{
	text := fmt.Sprintf("SBAR: %d \tBAR:%d \tBEAT:%d \t STEP:%d \t SSTEP:%d",superbar,bar, beat, beatstep, beatsubstep)
	return text
}


func showterminal() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	boldStyle := defStyle.Bold(true)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Event loop
	quit := func() {
		s.Fini()
		os.Exit(0)
	}


	seconds := time.Tick(time.Second/60)
	events := make(chan tcell.Event)
	go func() {
		for {
			events <- s.PollEvent()
		}
	}()

	for {

		select {

			case <-seconds:

			case event := <-events:
				switch ev:=event.(type) {
				case *tcell.EventResize:
				s.Sync()
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
						quit()
					} else if ev.Key() == tcell.KeyCtrlL {
						s.Sync()
					}

					switch(ev.Rune()){


						case ' ':
							samples[editsample].play()

						case 'z':
							pausedrums = !pausedrums
						case 'x':
							boosttones = !boosttones



						case 'f':
							frozentones=!frozentones

						case 't':
							toneall=!toneall


						case 'j':
							if(cursorpos<20){
								cursorpos++
							}

						case 'k':
							if (cursorpos>0){
								cursorpos--
							}

						case 'g':
							cursorpos=0

						case '0':
							screenpos=0
						case '1':
							screenpos=1
						case '2':
							screenpos=2
						case '3':
							screenpos=3
						case '4':
							screenpos=4
						case '5':
							screenpos=5
						case '6':
							screenpos=6
						case '7':
							screenpos=7
						case '8':
							screenpos=8
					}

					switch(screenpos){
						case 1:
							mainMenuControl(ev.Rune())
						case 2:
							phaseMenuControl(ev.Rune())
						case 3:
							drumMenuControl(ev.Rune())
						case 4:
							genMenuControl(ev.Rune())
						case 5:
							flexMenuControl(ev.Rune())
						case 6:
							audioMenuControl(ev.Rune())
						case 7:
							sampleMenuControl(ev.Rune())
						case 8:
							songMenuControl(ev.Rune())




					}


					limitParameters()
				}


		case <-autoregens:
			generatePatterns()


		}



		s.Clear()

		drawText(s, 0, 0, 40, 0, boldStyle, "TXTRTN - Terminal Extratone")
		drawText(s, 3, 3+cursorpos, 3, 3+cursorpos, defStyle, ">")

		switch(screenpos){
			case 0:
				drawText(s, 0, 1, 40, 1, boldStyle, "0. Welcome")
				helpDisplay(&s, defStyle)

			case 1:
				drawText(s, 0, 1, 40, 1, boldStyle, "1. General Settings")
				mainMenuDisplay(&s, defStyle)
			case 2:
				drawText(s, 0, 1, 40, 1, boldStyle, "2. Phase Settings")
				phaseMenuDisplay(&s, defStyle)
			case 3:
				drawText(s, 0, 1, 40, 1, boldStyle, "3. Drum Settings")
				drumMenuDisplay(&s, defStyle)
			case 4:
				drawText(s, 0, 1, 40, 1, boldStyle, "4. Generation Settings")
				genMenuDisplay(&s, defStyle)
			case 5:
				drawText(s, 0, 1, 40, 1, boldStyle, "5. Flex Settings")
				flexMenuDisplay(&s, defStyle)

			case 6:
				drawText(s, 0, 1, 40, 1, boldStyle, "6. Audio Settings")
				audioMenuDisplay(&s, defStyle)

			case 7:
				drawText(s, 0, 1, 40, 1, boldStyle, "7. Sampler Settings")
				sampleMenuDisplay(&s, defStyle)
			case 8:
				drawText(s, 0, 1, 40, 1, boldStyle, "8. Song Settings")
				songMenuDisplay(&s, defStyle)

		}


		drawText(s, 5, 20, 100, 20, defStyle, writeState())
		drawText(s, 5, 21, 50, 21, defStyle, writeTimings())
		drawText(s, 5, 22, 100, 22, defStyle, writePhases())
		drawText(s, 5, 23, 100, 23, defStyle, writeFlexing())


//		drawText(s, 10, 4, 50, 4, defStyle, fmt.Sprintf("Distortion: %v",distortion))


	//	drawText(s, 10, 20, 100, 25, defStyle, drumNames())
		// Update screen
		s.Show()

	}
}
