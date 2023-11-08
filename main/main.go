package main

import (
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "travel-booking-portal"
	"travel-booking-portal/config"
	"travel-booking-portal/gen/restapi"
	"travel-booking-portal/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)

	serverPort := viper.GetString(config.ServerPort)
	if serverPort == "" {
		serverPort = "8080"
	}

	port, err := strconv.Atoi(serverPort)
	if err != nil {
		panic(err)
	}

	server.Port = port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
