package main

import (
	"strconv"
	"time"

	"github.com/TBgo93/poc-goland-rbpi/textview"
	"github.com/TBgo93/poc-goland-rbpi/utils"
)

func main() {
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
		ip := "IP: " + utils.GetLocalIP().String()
		temp := utils.GetTempCore()

		cpu := utils.GetCPUInfo()
		var arrayTexts []string

		arrayTexts = append(arrayTexts, " ")
		arrayTexts = append(arrayTexts, ip)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, cpu)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, ram)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, temp)
		arrayTexts = append(arrayTexts, "---------------")

		tv.DrawListOfStrings(arrayTexts)

		time.Sleep(1 * time.Second)
	}
}
