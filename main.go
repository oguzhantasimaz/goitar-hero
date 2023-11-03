package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gdamore/tcell"
	"github.com/oguzhantasimaz/goitar-hero/models"
)

const (
	ALocation = 0
	SLocation = 4
	JLocation = 8
	KLocation = 12
	LLocation = 16
)

func main() {
	music, err := os.Open("scar_tissue.mp3")
	if err != nil {
		panic(err)
	}

	streamer, format, err := mp3.Decode(music)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	aNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	sNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	jNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	kNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue)
	lNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorOrange)

	ScarTissue := &models.Song{
		Notes: []*models.Note{
			{Time: 0, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 6, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 8, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 10, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 12, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 14, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 16, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 18, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 20, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 22, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 23, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 24, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 26, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 28, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 30, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 32, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 34, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 36, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 38, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 40, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 46, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 50, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 52, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 54, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 56, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 58, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 60, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 62, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 63, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 64, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 66, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 67, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 68, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
		},
	}

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	gameTime := 0
	score := 0

	go func() {
		<-done
		s.Fini()
		fmt.Println("Score: ", score)
		os.Exit(0)
	}()

	// Handle events
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					s.Fini()
					done <- true
				} else if ev.Rune() == 'a' || ev.Rune() == 'A' {
					checkNotePress(*ScarTissue, gameTime, "A", &score)
					//Simulate note press with * character
					drawText(s, ALocation, 8, 0, 0, aNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 's' || ev.Rune() == 'S' {
					checkNotePress(*ScarTissue, gameTime, "S", &score)
					//Simulate note press with * character
					drawText(s, SLocation, 8, 0, 0, sNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'j' || ev.Rune() == 'J' {
					checkNotePress(*ScarTissue, gameTime, "J", &score)
					//Simulate note press with * character
					drawText(s, JLocation, 8, 0, 0, jNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'k' || ev.Rune() == 'K' {
					checkNotePress(*ScarTissue, gameTime, "K", &score)
					//Simulate note press with * character
					drawText(s, KLocation, 8, 0, 0, kNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'l' || ev.Rune() == 'L' {
					checkNotePress(*ScarTissue, gameTime, "L", &score)
					//Simulate note press with * character
					drawText(s, LLocation, 8, 0, 0, kNoteStyle, "*")
					s.Show()
				}
			}
		}
	}()

	t := time.NewTicker(time.Second / 8)
	songStarted := false
	// Main game loop
	for {
		if gameTime > 3 && !songStarted {
			speaker.Play(beep.Seq(streamer, beep.Callback(func() {
				done <- true
			})))
			defer speaker.Close()
			songStarted = true
		}

		if gameTime > ScarTissue.Notes[len(ScarTissue.Notes)-1].Time+9 {
			done <- true
		}

		for _, n := range ScarTissue.Notes {
			if n.Time <= gameTime && n.Y < 9 {
				drawText(s, n.X, n.Y, 0, 0, n.Style, n.Key)
				n.Y = n.Y + 1
			}
		}

		drawText(s, 20, 0, 30, 0, aNoteStyle, fmt.Sprintf("Score: %d", score))
		s.Show()

		select {
		case <-t.C:
			gameTime++
			s.Clear()
			drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
			s.Sync()
			s.Show()
		}

	}

}

// checkNotePress: Checks if a note was pressed and updates the score
func checkNotePress(song models.Song, gameTime int, key string, score *int) {
	for _, n := range song.Notes {
		if n.Key == key && n.Y == 9 && gameTime == n.Time+8 {
			n.Y++
			*score = *score + 1
		}
	}
}

// drawText: Draws text on the screen
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

// drawNoteLocations: Draws the note locations on the screen
func drawNoteLocations(s tcell.Screen, styles []tcell.Style) {
	drawText(s, ALocation, 8, 0, 0, styles[0], "A")
	drawText(s, SLocation, 8, 0, 0, styles[1], "S")
	drawText(s, JLocation, 8, 0, 0, styles[2], "J")
	drawText(s, KLocation, 8, 0, 0, styles[3], "K")
	drawText(s, LLocation, 8, 0, 0, styles[4], "L")
}
