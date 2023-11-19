package main

import (
	"strconv"
	"time"

	"github.com/TBgo93/poc-goland-rbpi/textview"
	"github.com/TBgo93/poc-goland-rbpi/utils"
)

func main() {
	// Draw a text
	opts := textview.DefaultOpts
	tv := textview.NewWithOptions(opts)
	for {
		mem := utils.ReadMemoryStats()
		RAM := mem.MemTotal - mem.MemAvailable
		text := strconv.Itoa(RAM) + "%"

		tv.Draw("")
		tv.DrawChars("Stats: ")
		tv.DrawChars(text)

		time.Sleep(3 * time.Second)
	}
}
