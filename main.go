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
		freeRamPercentage := mem.MemAvailable * 100 / mem.MemTotal
		parsedFreeRamPercentage := strconv.Itoa(freeRamPercentage) + "%"

		// Funciona pero muestra una linea, se resetea
		// y vuelve a mostrar linea
		// tv.DrawChars("Stats: " + "\n" + parsedFreeRamPercentage)
		ram := "Uso de RAM: " + parsedFreeRamPercentage
		var arrayTexts []string

		arrayTexts = append(arrayTexts, "Stats: ")
		arrayTexts = append(arrayTexts, ram)

		tv.DrawListOfStrings(arrayTexts)

		time.Sleep(1 * time.Second)
	}
}
