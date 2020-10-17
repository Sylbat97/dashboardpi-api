package data

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetCPUData() (CPU, error) {
	var cpu CPU
	cmd := "mpstat 1 1 | tail -n 1 | awk '{print $12}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return cpu, errors.New("Cannot execute mpstat command")
	}
	result := strings.TrimSuffix(string(out), "\n")
	cpuUsage, err := strconv.ParseFloat(result, 64)
	cpuUsage = 100.00 - cpuUsage
	if err != nil {
		fmt.Println(err.Error())
		return cpu, errors.New("Cannot process mpstat result")
	}
	return CPU{Usage: cpuUsage}, nil
}
