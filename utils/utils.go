package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

func ReadMemoryStats() Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)
	res := Memory{}
	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value
		}
	}
	return res
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func GetTempCore() string {
	data, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}

	tempCommandResult := string(data)
	tempCommandResult = tempCommandResult[:len(tempCommandResult)-1]
	parsedTemp, err := strconv.Atoi(tempCommandResult)

	if err != nil {
		panic(err)
	}

	calculatedTemp := int(parsedTemp / 1000)

	return fmt.Sprintf("Temp: %d °C", calculatedTemp)
}

func GetCPUInfo() string {
	percentageCpu, err := cpu.Percent(0, false)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Uso de CPU: %.2f%s", percentageCpu[0], "%")
}

// Revisar necesidad ---
func parseLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}
