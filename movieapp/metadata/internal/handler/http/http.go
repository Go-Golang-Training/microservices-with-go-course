package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller/metadata"
)

// Handler demonstrates the adapter pattern for HTTP transport
type Handler struct {
	ctrl *metadata.Controller
}

// New follows the Go convention of constructor functions
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMetadata handles HTTP requests - separation of transport concerns
func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		// Proper HTTP status code usage
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Context propagation for request-scoped data
	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)

	switch {
	case errors.Is(err, metadata.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Proper content-type setting
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
