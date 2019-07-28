package logic

import (
	"errors"

	"github.com/99designs/gqlgen/qfield"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/gendata"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/models"
)

func QueryDevices(fields qfield.QField) ([]*models.Device, error) {
	var result []*models.Device
	var err error
	for _, i := range gendata.Groups {
		for _, o := range i.Devices {
			device, err := QueryDevice(fields, o.ID)
			if err != nil {
				return result, err
			}
			result = append(result, device)
		}
	}
	return result, err
}

func QueryDevice(fields qfield.QField, id string) (*models.Device, error) {
	var result models.Device
	var err error
	r, exists := gendata.GetDevice(id)
	if !exists {
		return &result, errors.New("Device does not exist")
	}
	result = models.CastDevice(r)
	if fields.ContainsDirectChild("group") {
		_, groupFields := fields.GetDirectChild("group")
		result.Group, err = QueryGroup(groupFields, r.Group)
	}

	return &result, err
}
