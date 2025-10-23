package main

import (
	"context"
	"production_service/bootstrap"
	"production_service/infrastructure/grpc_client"
	"production_service/infrastructure/grpc_service"
	harvest_record_service "production_service/infrastructure/grpc_service/harvest_record"
	pest_disease_record_service "production_service/infrastructure/grpc_service/pest_disease_record"

	"github.com/anhvanhoa/service-core/domain/discovery"
	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
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

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

	harvestRecordService := harvest_record_service.NewHarvestRecordService(app.Repos)
	pestDiseaseRecordService := pest_disease_record_service.NewPestDiseaseRecordService(app.Repos)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		app.Cache,
		harvestRecordService,
		pestDiseaseRecordService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
