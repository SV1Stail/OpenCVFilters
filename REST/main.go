package main

import (
	"net/http"
	"os"

	"github.com/SV1Stail/OpenCVFilters/REST/constants"
	"github.com/SV1Stail/OpenCVFilters/REST/gen"
	httpserver "github.com/SV1Stail/OpenCVFilters/REST/http_server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	RestPort          string
	GrpcServerAddress string
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	RestPort = os.Getenv("REST_CONTAINER_PORT")
	GrpcServerAddress = os.Getenv("GRPC_SERVER_ADDRESS")

}

func main() {

	if GrpcServerAddress == "" {
		log.Err(constants.ErrBadRequest).Msg("empty grpc server address")
		GrpcServerAddress = "localhost:50051"
	}

	conn, err := grpc.NewClient(
		GrpcServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Err(err).Msg("Failed to connect to gRPC server")
		os.Exit(1)
	}
	client := httpserver.NewClient(gen.NewServiceClient(conn))

	http.HandleFunc("/upload", client.UploadHandler)
	fs := http.FileServer(http.Dir("/web_ui"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "/web_ui/ui.html")
			return
		}
		fs.ServeHTTP(w, r)
	})

	log.Info().Msgf("HTTP server running on :%s...", RestPort)
	err = http.ListenAndServe(":"+RestPort, nil)
	log.Err(err).Msg("HTTP server fall")

}
