package grpc_service

import (
	"production_service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	harvestRecordService proto_harvest_record.HarvestRecordServiceServer,
	pestDiseaseRecordService proto_pest_disease_record.PestDiseaseRecordServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_harvest_record.RegisterHarvestRecordServiceServer(server, harvestRecordService)
			proto_pest_disease_record.RegisterPestDiseaseRecordServiceServer(server, pestDiseaseRecordService)
		},
	)
}
