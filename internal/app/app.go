package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vgekko/anistream-gateway/config"
	"github.com/vgekko/anistream-gateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net/http"

	"github.com/vgekko/anistream-gateway/internal/service/content/pb"
)

func RunGateway(cfg *config.Config) {
	log := logger.New(cfg.Env)

	grpcGatewayMux := runtime.NewServeMux()

	connContent, err := grpc.Dial(cfg.Microservice.ContentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("failed to try dial", slog.String("port", cfg.Microservice.ContentAddr))
	}
	defer connContent.Close()

	if err := pb.RegisterContentServiceHandler(context.Background(), grpcGatewayMux, connContent); err != nil {
		log.Error("failed to initialize handler", slog.String("service", "content"))
	}

	httpMux := http.NewServeMux()

	httpMux.Handle("/v1/search", grpcGatewayMux)

	log.Info("starting grpc-gateway")

	if err := http.ListenAndServe(cfg.ProxyAddr, httpMux); err != nil {
		log.Error("failed to server", slog.String("addr", cfg.ProxyAddr))
	}
}
