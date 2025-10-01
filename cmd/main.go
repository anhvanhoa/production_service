package main

import (
	"context"
	"production_service/bootstrap"
	"production_service/infrastructure/grpc_service"
	harvest_record_service "production_service/infrastructure/grpc_service/harvest_record"
	pest_disease_record_service "production_service/infrastructure/grpc_service/pest_disease_record"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	harvestRecordService := harvest_record_service.NewHarvestRecordService(app.Repos)
	pestDiseaseRecordService := pest_disease_record_service.NewPestDiseaseRecordService(app.Repos)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		harvestRecordService,
		pestDiseaseRecordService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
