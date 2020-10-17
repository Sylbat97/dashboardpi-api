package data

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

//CPU information
type CPU struct {
	Usage float64 `json:"Usage"`
	Temp  float64 `json:"Temp"`
}

//GetCPUData returns general informations about CPU usage
func GetCPUData() (CPU, error) {
	var cpu CPU
	var errstr string
	var wg sync.WaitGroup
	wg.Add(2)
	go getUsageData(&wg, &cpu, &errstr)
	go getTempData(&wg, &cpu, &errstr)
	wg.Wait()
	if errstr != "" {
		return cpu, errors.New(errstr)
	}

	return cpu, nil
}

func getUsageData(wg *sync.WaitGroup, cpu *CPU, errstr *string) {
	fmt.Println("Getting usage")
	cmd := "mpstat 1 1 | tail -n 1 | awk '{print $12}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		wg.Done()
		*errstr = err.Error()
		return
	}
	result := strings.ReplaceAll(strings.TrimSuffix(string(out), "\n"), ",", "")
	cpuUsage, err := strconv.ParseFloat(result, 64)
	fmt.Println(cpuUsage)
	cpuUsage = 100.00 - cpuUsage
	if err != nil {
		wg.Done()
		*errstr = err.Error()
		return
	}
	cpu.Usage = cpuUsage
	wg.Done()
}

func getTempData(wg *sync.WaitGroup, cpu *CPU, errstr *string) {
	fmt.Println("Getting temp")
	cmd := "vcgencmd measure_temp | cut -b 6- | rev | cut -c 3- | rev"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		wg.Done()
		*errstr = err.Error()
		return
	}
	result := strings.ReplaceAll(strings.TrimSuffix(string(out), "\n"), ",", "")
	cpuTemp, err := strconv.ParseFloat(result, 64)
	if err != nil {
		wg.Done()
		*errstr = err.Error()
		return
	}
	cpu.Temp = cpuTemp
	wg.Done()
}
