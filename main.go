package main

import (
	"image/color"
	"strconv"
	"time"

	"github.com/TBgo93/poc-goland-rbpi/display"
	"github.com/TBgo93/poc-goland-rbpi/textview"
	"github.com/TBgo93/poc-goland-rbpi/utils"
)

func main() {
	dsp, err := display.Init()
	if err != nil {
		panic(err)
	}
	defer dsp.Close()

	dsp.Rotate(display.ROTATION_90)
	// Set the screen color to white
	dsp.FillScreen(color.RGBA{R: 0, G: 0, B: 0, A: 0})

	// Draw a text
	opts := textview.DefaultOpts
	tv := textview.NewWithOptions(opts)
	for true == true {
		mem := utils.ReadMemoryStats()
		RAM := mem.MemTotal - mem.MemAvailable
		text := strconv.Itoa(RAM) + "%"

		tv.Draw(text)
		time.Sleep(3 * time.Second)
	}
}
