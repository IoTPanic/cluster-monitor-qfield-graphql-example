package ql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/qfield"

	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/logic"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Devices(ctx context.Context) ([]*models.Device, error) {
	qfields := qfield.GetQField(ctx, graphql.GetRequestContext(ctx).Variables)
	return logic.QueryDevices(qfields)
}
func (r *queryResolver) Groups(ctx context.Context) ([]*models.DeviceGroup, error) {
	qfields := qfield.GetQField(ctx, graphql.GetRequestContext(ctx).Variables)
	return logic.QueryGroups(qfields)
}
func (r *queryResolver) Installations(ctx context.Context) ([]*models.Installation, error) {
	qfields := qfield.GetQField(ctx, graphql.GetRequestContext(ctx).Variables)
	return logic.QueryInstallations(qfields)
}
func (r *queryResolver) Device(ctx context.Context, id string) (*models.Device, error) {
	qfields := qfield.GetQField(ctx, graphql.GetRequestContext(ctx).Variables)
	return logic.QueryDevice(qfields, id)
}
