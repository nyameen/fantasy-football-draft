package draftclock

import (
	"fmt"
	"sync/atomic"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DraftClock struct {
	ClockObjects fyne.CanvasObject

	clock     *widget.RichText
	min       int
	sec       int
	pause     atomic.Bool
	isRunning atomic.Bool
	restart   chan int
}

// NewDraftClock create a new Draft Clock object to handle timer countdown
func NewDraftClock() (*DraftClock, error) {
	dc := &DraftClock{
		pause:     atomic.Bool{},
		isRunning: atomic.Bool{},
		restart:   make(chan int, 1),
		clock: widget.NewRichText(&widget.TextSegment{
			Text: "",
			Style: widget.RichTextStyle{
				Alignment: fyne.TextAlignCenter,
				SizeName:  theme.SizeNameHeadingText,
			},
		}),
	}

	// Each person has 3 mins to draft
	dur, err := time.ParseDuration("3m")
	if err != nil {
		return nil, err
	}

	// get the number of minutes and seconds to display
	dc.min = int(dur.Seconds() / 60)
	dc.sec = int(int(dur.Seconds()) % 60)
	dc.setDefaultClock()

	// create three buttons to Start, Pause, and Restart
	startButton := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), dc.onPlayTap)
	pauseButton := widget.NewButtonWithIcon("Pause", theme.MediaPauseIcon(), dc.onPauseTap)
	restartButton := widget.NewButtonWithIcon("Restart", theme.MediaReplayIcon(), dc.onRestartTap)

	// Bundle it into a grid
	dc.ClockObjects = container.NewGridWithColumns(4, dc.clock, startButton, pauseButton, restartButton)
	return dc, nil
}

func (d *DraftClock) onPlayTap() {
	// If it is paused, unpause
	d.pause.Store(false)

	// if a countdown is not currently going, start one
	if !d.isRunning.Load() {
		d.isRunning.Store(true)
		go d.countdown()
	}
}

func (d *DraftClock) onPauseTap() {
	d.pause.Store(true)
}

func (d *DraftClock) onRestartTap() {
	// pause the countdown
	d.pause.Store(true)

	// if a countdown is currently running, stop it and make it return
	if d.isRunning.Load() {
		d.isRunning.Store(false)
		d.restart <- 1
	}

	// set the time to the default 3 mins
	d.setDefaultClock()
}

func (d *DraftClock) setDefaultClock() {
	d.clock.Segments = []widget.RichTextSegment{&widget.TextSegment{
		Text: fmt.Sprintf("%.2d:%.2d\n", d.min, d.sec),
		Style: widget.RichTextStyle{
			Alignment: fyne.TextAlignCenter,
			SizeName:  theme.SizeNameHeadingText,
		},
	}}
	d.clock.Refresh()
}

func (d *DraftClock) countdown() {
	min := d.min
	sec := d.sec
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		// every second decrement a sec and display
		case <-ticker.C:
			// pause was pushed, don't decrement
			if d.pause.Load() {
				continue
			}

			// timer has ended restore defaults
			if min == 0 && sec == 0 {
				d.isRunning.Store(false)
				d.setDefaultClock()
				return
			}

			// seconds go from 0 - 59
			if sec == 00 {
				sec = 59
			} else {
				sec -= 1
			}

			// can't have negative minutes
			if sec == 59 && min != 0 {
				min -= 1
			}

			// display a new time
			d.clock.Segments = []widget.RichTextSegment{&widget.TextSegment{
				Text: fmt.Sprintf("%.2d:%.2d\n", min, sec),
				Style: widget.RichTextStyle{
					Alignment: fyne.TextAlignCenter,
					SizeName:  theme.SizeNameHeadingText,
				},
			}}
			d.clock.Refresh()

		// restart button was pressed, stop counting down
		case <-d.restart:
			return
		}
	}
}
