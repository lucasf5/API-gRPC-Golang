package main

import (
	"grpcTreino/src/api/models"
	"grpcTreino/src/api/routes"
)

func main() {
	routes.HandleRequests()
	models.Connection()
}
