package main

import (
	"github.com/cankirma/go-hexagonal-architecture-example/internal/adapters/app/api"
	"github.com/cankirma/go-hexagonal-architecture-example/internal/adapters/core/arithmetic"
	gRPC "github.com/cankirma/go-hexagonal-architecture-example/internal/adapters/framework/left/grpc"
	"github.com/cankirma/go-hexagonal-architecture-example/internal/adapters/framework/right/db"
	"github.com/cankirma/go-hexagonal-architecture-example/internal/ports"
	"log"
	"os"
)

func main() {
	var err error
	var dbasAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")
	dbasAdapter, err = db.NewAdapter(dbaseDriver, dSourceName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer dbasAdapter.CloseDbConnection()
	core = arithmetic.NewArith()
	appAdapter = api.NewAdapter(dbasAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}

//username:password@protocol(address)/dbname?param=value
