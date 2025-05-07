package guest

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"net/http"
)

type Handler struct {
	log    *slog.Logger
	ctx    context.Context
	client *mongo.Client
}

func NewGuestHandler(ctx context.Context, logger *slog.Logger, client *mongo.Client) *Handler {
	return &Handler{
		log:    logger,
		ctx:    ctx,
		client: client,
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.log.Debug("Request Received", slog.String("Path", req.URL.Path), slog.String("Method", req.Method))
	if req.Method == http.MethodPost {
		handlePost(rw, req, h.client)
		return
	} else {
		http.Error(rw, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}
