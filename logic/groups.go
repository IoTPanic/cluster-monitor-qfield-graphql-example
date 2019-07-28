package logic

import (
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/qfield"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/gendata"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/models"
)

func QueryGroups(fields qfield.QField) ([]*models.DeviceGroup, error) {
	var result []*models.DeviceGroup
	for _, i := range gendata.Groups {
		g, err := QueryGroup(fields, i.ID)
		if err != nil {
			return result, err
		}
		result = append(result, g)
	}
	return result, nil
}

func QueryGroup(fields qfield.QField, id string) (*models.DeviceGroup, error) {
	var result models.DeviceGroup
	var err error
	g, exists := gendata.GetGroup(id)
	if !exists {
		return &result, errors.New("Group does not exits")
	}
	result = models.CastGroup(g)
	if fields.ContainsDirectChild("installation") {
		_, installationFields := fields.GetDirectChild("installation")
		iid, exists := gendata.GetInstallationIDForGroup(g.ID)
		if !exists {
			return &result, errors.New("Installation does not exist")
		}
		result.Installation, err = QueryInstallation(installationFields, iid)
	}
	if fields.ContainsDirectChild("devices") {
		fmt.Println("GCD")
		_, deviceFields := fields.GetDirectChild("devices")
		for _, i := range g.Devices {
			d, err := QueryDevice(deviceFields, i.ID)
			if err != nil {
				return &result, errors.New("Device does not exist")
			}
			result.Devices = append(result.Devices, d)
		}
	}
	return &result, err
}
