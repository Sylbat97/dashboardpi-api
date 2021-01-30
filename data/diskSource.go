package data

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

//Disk information
type Disk struct {
	Partitions []Part `json:"Partitions"`
}

//Sizes in KB
type Part struct {
	FileSystem string  `json:"FileSystem"`
	Type       string  `json:"Type"`
	Size       float64 `json:"Size"`
	Used       float64 `json:"Used"`
	Avail      float64 `json:"Avail"`
	Percentage float64 `json:"Percentage"`
	MountPoint string  `json:"MountPoint"`
}

//GetDiskData returns general informations about CPU usage
func GetDiskData() (Disk, error) {
	var disk Disk
	var errstr string

	fmt.Println("Getting disk usage")
	//Sizes in kB
	cmd := "df -x tmpfs -x devtmpfs -Tk | tail -n +2"

	out, err := exec.Command("bash", "-c", cmd).Output()
	result := string(out)

	if err != nil {
		errstr = err.Error()
		return disk, errors.New(errstr)
	}

	scanner := bufio.NewScanner(strings.NewReader(result))
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)
		var partition Part
		partition.FileSystem = data[0]
		partition.Type = data[1]
		size, _ := strconv.ParseFloat(data[2], 64)
		used, _ := strconv.ParseFloat(data[3], 64)
		avail, _ := strconv.ParseFloat(data[4], 64)
		partition.Size = size
		partition.Used = used
		partition.Avail = avail
		partition.Percentage = (used / size) * 100
		partition.MountPoint = data[6]
		disk.Partitions = append(disk.Partitions, partition)
	}

	return disk, nil
}
