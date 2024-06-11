package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	cpuTicks  = 100
	cpuFields = 8
	cpuMax    = 1000
	statFile  = "/proc/stat"
)

type (
	textReadOptions struct {
		keepSpace     bool
		withoutBlanks bool
		omitPrefix    string
	}

	// TextReadOption defines the method to customize the text reading functions.
	TextReadOption func(*textReadOptions)
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

func SystemCpuUsage() (uint64, error) {
	lines, err := readTextLines(statFile, withoutBlank())
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			if len(fields) < cpuFields {
				return 0, fmt.Errorf("bad format of cpu stats")
			}

			var totalClockTicks uint64
			for _, i := range fields[1:cpuFields] {
				v, err := parseUint(i)
				if err != nil {
					return 0, err
				}

				totalClockTicks += v
			}

			return (totalClockTicks * uint64(time.Second)) / cpuTicks, nil
		}
	}

	return 0, errors.New("bad stats format")
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

func parseUint(s string) (uint64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			return 0, nil
		}

		return 0, fmt.Errorf("cgroup: bad int format: %s", s)
	}

	if v < 0 {
		return 0, nil
	}

	return uint64(v), nil
}

func readTextLines(filename string, opts ...TextReadOption) ([]string, error) {
	var readOpts textReadOptions
	for _, opt := range opts {
		opt(&readOpts)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !readOpts.keepSpace {
			line = strings.TrimSpace(line)
		}
		if readOpts.withoutBlanks && len(line) == 0 {
			continue
		}
		if len(readOpts.omitPrefix) > 0 && strings.HasPrefix(line, readOpts.omitPrefix) {
			continue
		}

		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func withoutBlank() TextReadOption {
	return func(o *textReadOptions) {
		o.withoutBlanks = true
	}
}
