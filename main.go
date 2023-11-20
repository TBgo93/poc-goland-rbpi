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
		freeRamPercentage := mem.MemFree * 100 / mem.MemTotal
		parsedFreeRamPercentage := strconv.Itoa(freeRamPercentage) + "%"

		// Funciona pero muestra una linea, se resetea
		// y vuelve a mostrar linea
		tv.DrawChars("Stats: ")
		tv.DrawChars(parsedFreeRamPercentage)

		time.Sleep(3 * time.Second)
	}
}
