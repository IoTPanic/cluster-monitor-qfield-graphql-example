package gendata

import "fmt"

type Device struct {
	ID            string
	Name          string
	Group         string
	OS            string
	State         string
	CPUPrecentage int
	MemoryMb      int
}

type DeviceGroup struct {
	ID      string
	Name    string
	Devices map[string]Device
}

type Installation struct {
	ID       string
	Name     string
	Location string
	Groups   []string
}

var Installations []Installation
var Groups []DeviceGroup

func CreateMockData() {
	Installations = append(Installations, Installation{
		ID:       "1",
		Name:     "Test Installation 1",
		Location: "Saco, ME",
		Groups:   make([]string, 2),
	})
	Installations[0].Groups[0] = "1"
	Installations[0].Groups[1] = "2"
	Groups = append(Groups, DeviceGroup{
		ID:   "1",
		Name: "Rack 1",
		Devices: map[string]Device{
			"1": Device{
				ID:            "1",
				Name:          "Server 1",
				Group:         "1",
				OS:            "Linux Arch",
				State:         "Online",
				CPUPrecentage: 23,
				MemoryMb:      4096,
			},
			"2": Device{
				ID:            "2",
				Name:          "Server 2",
				Group:         "1",
				OS:            "Linux CentOS",
				State:         "Rebooting",
				CPUPrecentage: 0,
				MemoryMb:      0,
			},
		},
	})
	Groups = append(Groups, DeviceGroup{
		ID:   "2",
		Name: "Office 1",
		Devices: map[string]Device{
			"1": Device{
				ID:            "3",
				Name:          "Chromecase",
				Group:         "2",
				OS:            "Linux Chrome",
				State:         "Online",
				CPUPrecentage: 23,
				MemoryMb:      0,
			},
			"2": Device{
				ID:            "4",
				Name:          "Stacie's Laptop",
				Group:         "2",
				OS:            "Windows 10",
				State:         "Offline",
				CPUPrecentage: 0,
				MemoryMb:      0,
			},
		},
	})
	fmt.Println(Installations)
	fmt.Println(Groups)
}

func GetDevice(id string) (Device, bool) {
	for _, i := range Groups {
		for _, o := range i.Devices {
			if o.ID == id {
				return o, true
			}
		}
	}
	return Device{}, false
}

func GetGroup(id string) (DeviceGroup, bool) {
	for _, i := range Groups {
		if i.ID == id {
			return i, true
		}
	}
	return DeviceGroup{}, false
}

func GetInstallation(id string) (Installation, bool) {
	for _, i := range Installations {
		if i.ID == id {
			return i, true
		}
	}
	return Installation{}, false
}

func GetInstallationIDForGroup(groupId string) (string, bool) {
	for _, i := range Installations {
		for _, o := range i.Groups {
			for _, p := range Groups {
				if o == p.ID {
					return i.ID, true
				}
			}
		}
	}
	return "", false
}

func GetDevicesFromInstallaion(installationId string) []string {
	var result []string
	i, exists := GetInstallation(installationId)
	if !exists {
		return []string{}
	}
	for _, i := range i.Groups {
		g, exists := GetGroup(i)
		if !exists {
			continue
		}
		for _, o := range g.Devices {
			result = append(result, o.ID)
		}
	}
	return result
}
