# GQLGen GraphQL Library QField Example - Cluster Monitor

This repo is an example simple GraphQL API that uses QFIelds to find nested fields within a query which allows for true and easy graphQL functionality. All code is written in golang, and uses the `dep` tool to install all dependancies.

To Intall Depandancies- `dep ensure`

To Run- `go run main.go`

The schema and an interface to interact can be found in the playground at `localhost:8080/query/playground`

Please note the dataset `gendata` is simply meant to demonstrate the QFields, I know it's ineffiecient and you wouldn't want to query a dataset like that.