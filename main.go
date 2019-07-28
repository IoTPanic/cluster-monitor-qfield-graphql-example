package main

import (
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/api"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/gendata"
)

func main() {
	gendata.CreateMockData()
	api.Start(":8080")
}
