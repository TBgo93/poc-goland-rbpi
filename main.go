package main

import (
	"fmt"
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
	dsp.FillScreen(color.RGBA{R: 255, G: 255, B: 255, A: 0})

	// Draw a text
	opts := textview.DefaultOpts
	tv := textview.NewWithOptions(opts)
	for {
		mem := utils.ReadMemoryStats()
		RAM := mem.MemTotal - mem.MemAvailable
		text := strconv.Itoa(RAM) + "%"
		fmt.Println(text)

		tv.Draw(text)
		time.Sleep(3 * time.Second)
	}
}
