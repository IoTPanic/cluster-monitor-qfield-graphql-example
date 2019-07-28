package logic

import (
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/qfield"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/gendata"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/models"
)

func QueryInstallations(fields qfield.QField) ([]*models.Installation, error) {
	var result []*models.Installation
	for _, i := range gendata.Installations {
		in, err := QueryInstallation(fields, i.ID)
		if err != nil {
			return result, err
		}
		result = append(result, in)
	}
	return result, nil
}

func QueryInstallation(fields qfield.QField, id string) (*models.Installation, error) {
	var result models.Installation
	i, exists := gendata.GetInstallation(id)
	if !exists {
		return &result, errors.New("Installation does not exits")
	}
	result = models.CastInstallation(i)
	if fields.ContainsDirectChild("groups") {
		_, groupFields := fields.GetDirectChild("groups")
		for _, i := range i.Groups {
			g, err := QueryGroup(groupFields, i)
			if err != nil {
				return &result, err
			}
			result.Groups = append(result.Groups, g)
		}
	}
	if fields.ContainsDirectChild("devices") {
		_, deviceFields := fields.GetDirectChild("devices")
		devices := gendata.GetDevicesFromInstallaion(id)
		for _, i := range devices {
			d, err := QueryDevice(deviceFields, i)
			if err != nil {
				return &result, errors.New("Devices does not exist")
			}
			result.Devices = append(result.Devices, d)
		}
	}
	fmt.Println(result.Groups[0].Devices[0].Group)
	return &result, nil
}
