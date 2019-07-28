package models

import (
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/gendata"
)

func CastDevice(d gendata.Device) Device {
	status := Status{
		State: d.State,
	}
	if d.CPUPrecentage == 0 {
		status.CPU = nil
	} else {
		status.CPU = &d.CPUPrecentage
	}
	if d.MemoryMb == 0 {
		status.MemoryMb = nil
	} else {
		status.MemoryMb = &d.MemoryMb
	}
	return Device{
		ID:     d.ID,
		Name:   d.Name,
		Status: &status,
		Os:     d.OS,
	}
}

func CastGroup(g gendata.DeviceGroup) DeviceGroup {
	result := DeviceGroup{
		ID:   g.ID,
		Name: g.Name,
	}

	for _, i := range g.Devices {
		d := CastDevice(i)
		result.Devices = append(result.Devices, &d)
	}
	return result
}

func CastInstallation(i gendata.Installation) Installation {
	return Installation{
		ID:       i.ID,
		Location: i.Location,
		Name:     i.Name,
	}
}
