package guest

import (
	"context"
	"log/slog"
	"net/http"
)

type Handler struct {
	log *slog.Logger
	ctx context.Context
}

func NewGuestHandler(ctx context.Context, logger *slog.Logger) *Handler {
	return &Handler{
		log: logger,
		ctx: ctx,
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.log.Debug("Request Received", slog.String("Path", req.URL.Path), slog.String("Method", req.Method))
	if req.Method == http.MethodPost {
		handlePost(rw, req)
		return
	} else {
		http.Error(rw, "invalid method", http.StatusMethodNotAllowed)
		return
	}
}
