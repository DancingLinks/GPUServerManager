package connect

import (
	"strconv"
	"strings"
)

type GPUList struct {
	ID string
	GPUStatus []GPUStatus

}

type GPUStatus struct {
	Ok 			bool
	Power		int
	Temperature int
	Utilization int
	Memory		int
}

func Parse(status string) GPUList {
	gpuList := GPUList{}
	split := strings.Split(status, "|")
	gpuList.ID = split[0]
	gpuList.GPUStatus = []GPUStatus{}
	split = split[1:len(split)-1]
	for i := range split {
		list := strings.Split(split[i], ",")
		status := GPUStatus{}
		if list[0] == "1" {
			status.Ok = true
			status.Power, _ = strconv.Atoi(list[1])
			status.Temperature, _ = strconv.Atoi(list[2])
			status.Utilization, _ = strconv.Atoi(list[3])
			status.Memory, _ = strconv.Atoi(list[4])
		} else {
			status.Ok = false
		}
		gpuList.GPUStatus = append(gpuList.GPUStatus, status)
	}
	return gpuList
}
