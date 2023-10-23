package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/oguzhantasimaz/goitar-hero/models"
	"log"
	"os"
	"time"
)

const (
	ALocation = 0
	SLocation = 4
	JLocation = 8
	KLocation = 12
	LLocation = 16
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	aNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	sNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	jNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	kNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue)
	lNoteStyle := tcell.StyleDefault.Foreground(tcell.ColorOrange)

	ExampleSong := &models.Song{
		Notes: []*models.Note{
			{Time: 0, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 4, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 5, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 6, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 10, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 11, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 12, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 13, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 20, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 21, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 23, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 25, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 27, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 30, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 34, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 36, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
			{Time: 38, Key: "A", X: ALocation, Y: 0, Style: aNoteStyle},
			{Time: 39, Key: "S", X: SLocation, Y: 0, Style: sNoteStyle},
			{Time: 40, Key: "J", X: JLocation, Y: 0, Style: jNoteStyle},
			{Time: 41, Key: "K", X: KLocation, Y: 0, Style: kNoteStyle},
			{Time: 42, Key: "L", X: LLocation, Y: 0, Style: lNoteStyle},
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
	events := make(chan tcell.Event)
	go func() {
		for {
			ev := s.PollEvent()
			events <- ev
		}
	}()

	go func() {
		for {
			ev := <-events
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					s.Fini()
					os.Exit(0)
				} else if ev.Rune() == 'a' || ev.Rune() == 'A' {
					checkNotePress(*ExampleSong, "A", &score)
					drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
					drawText(s, ALocation, 8, 0, 0, aNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 's' || ev.Rune() == 'S' {
					checkNotePress(*ExampleSong, "S", &score)
					drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
					drawText(s, SLocation, 8, 0, 0, sNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'j' || ev.Rune() == 'J' {
					checkNotePress(*ExampleSong, "J", &score)
					drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
					drawText(s, JLocation, 8, 0, 0, jNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'k' || ev.Rune() == 'K' {
					checkNotePress(*ExampleSong, "K", &score)
					drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
					drawText(s, KLocation, 8, 0, 0, kNoteStyle, "*")
					s.Show()
				} else if ev.Rune() == 'l' || ev.Rune() == 'L' {
					checkNotePress(*ExampleSong, "L", &score)
					drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
					drawText(s, LLocation, 8, 0, 0, kNoteStyle, "*")
					s.Show()
				}
			}
		}
	}()

	t := time.NewTicker(time.Second / 3)
	for {
		gameTime++

		if gameTime > ExampleSong.Notes[len(ExampleSong.Notes)-1].Time+9 {
			s.Fini()
			os.Exit(score)
		}

		for _, n := range ExampleSong.Notes {
			if n.Time <= gameTime && n.Y < 9 {
				drawText(s, n.X, n.Y, 0, 0, n.Style, n.Key)
				n.Y = n.Y + 1
			}
		}
		drawText(s, 20, 0, 30, 0, aNoteStyle, fmt.Sprintf("Score: %d", score))
		s.Show()
		select {
		case <-t.C:
			s.Clear()
			drawNoteLocations(s, []tcell.Style{aNoteStyle, sNoteStyle, jNoteStyle, kNoteStyle, lNoteStyle})
			s.Sync()
			s.Show()
		}

	}

}

func checkNotePress(song models.Song, key string, score *int) {
	for _, n := range song.Notes {
		if n.Key == key && n.Y == 9 {
			*score = *score + 1
		}
	}
}

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

func drawNoteLocations(s tcell.Screen, styles []tcell.Style) {
	drawText(s, ALocation, 8, 0, 0, styles[0], "A")
	drawText(s, SLocation, 8, 0, 0, styles[1], "S")
	drawText(s, JLocation, 8, 0, 0, styles[2], "J")
	drawText(s, KLocation, 8, 0, 0, styles[3], "K")
	drawText(s, LLocation, 8, 0, 0, styles[4], "L")
}
