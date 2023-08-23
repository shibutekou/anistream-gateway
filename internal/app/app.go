package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vgekko/anistream-gateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"net/http"

	"github.com/vgekko/anistream-gateway/internal/service/content/pb"
)

func RunHttpProxy(cfg *config.Config) {
	grpcGatewayMux := runtime.NewServeMux()

	connContent, err := grpc.Dial(cfg.ContentServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer connContent.Close()

	if err := pb.RegisterContentServiceHandler(context.Background(), grpcGatewayMux, connContent); err != nil {
		log.Fatal(err)
	}

	httpMux := http.NewServeMux()

	httpMux.Handle("/v1/search", grpcGatewayMux)

	slog.Info("starting HTTP server")

	log.Fatal(http.ListenAndServe(cfg.ProxyAddr, httpMux))
}
