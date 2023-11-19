package main

import (
	"fmt"
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
		fmt.Println(text)

		tv.Draw(text)
		time.Sleep(3 * time.Second)
	}
}
