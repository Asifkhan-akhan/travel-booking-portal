package main

import (
	"fmt"
	"runtime"
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/viper"
	runTime "travel-booking-portal"
	"travel-booking-portal/config"
	"travel-booking-portal/gen/restapi"
	"travel-booking-portal/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runTime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	//defer server.Shutdown()
	defer func() {
		if err := server.Shutdown(); err != nil {
			// Handle the error, e.g., log it or take other appropriate action.
			panic(err)
		}
	}()

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

	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of active goroutines: %d\n", numGoroutines)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
