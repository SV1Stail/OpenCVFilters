package main

import (
	"net"
	"os"

	"github.com/SV1Stail/OpenCVFilters/grpc/gen"
	"github.com/SV1Stail/OpenCVFilters/grpc/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	GrpcPort string
)

func Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	GrpcPort = os.Getenv("GRPC_PORT")
}

func main() {
	lis, err := net.Listen("tcp", ":"+GrpcPort)
	if err != nil {
		log.Err(err).Msg("failed start listen")
	}
	grpcServer := grpc.NewServer()
	gen.RegisterServiceServer(grpcServer, &server.Server{})

	log.Info().Msgf("Server is running on port %s...", GrpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Err(err).Msg("failed to serve")
	}
}
