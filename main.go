package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/TBgo93/poc-goland-rbpi/textview"
	"github.com/TBgo93/poc-goland-rbpi/utils"
)

func handleSigInt(displayText *textview.TextView) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for range c {
		fmt.Println("")
		fmt.Println("Exiting...")
		// off display
		var arrayTexts []string
		displayText.DrawListOfStrings(arrayTexts)
		os.Exit(0)
	}
}

func main() {
	// Init display
	opts := textview.DefaultOpts
	tv := textview.NewWithOptions(opts)

	handleSigInt(tv)

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
		countCoreLogical, countCorePhysical := utils.GetCores()

		var arrayTexts []string

		// arrayTexts = append(arrayTexts, " ")
		arrayTexts = append(arrayTexts, ip)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, cpu)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, countCoreLogical)
		arrayTexts = append(arrayTexts, countCorePhysical)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, ram)
		arrayTexts = append(arrayTexts, "---------------")
		arrayTexts = append(arrayTexts, temp)
		arrayTexts = append(arrayTexts, "---------------")

		tv.DrawListOfStrings(arrayTexts)

		time.Sleep(1 * time.Second)
	}
}
