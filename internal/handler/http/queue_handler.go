package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"queue-broker/internal/domain"
)

type QueueHandler struct {
	service        domain.QueueService
	defaultTimeout time.Duration
}

func NewQueueHandler(service domain.QueueService, defaultTimeout time.Duration) *QueueHandler {
	return &QueueHandler{
		service:        service,
		defaultTimeout: defaultTimeout,
	}
}

func (h *QueueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/queue/")
	if path == "" {
		http.Error(w, "Queue name required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		h.putMessage(w, r, path)
	case http.MethodGet:
		h.getMessage(w, r, path)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *QueueHandler) putMessage(w http.ResponseWriter, r *http.Request, queueName string) {
	var msg struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil || msg.Message == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.service.PutMessage(queueName, msg.Message); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *QueueHandler) getMessage(w http.ResponseWriter, r *http.Request, queueName string) {
	timeout := h.defaultTimeout
	if timeoutParam := r.URL.Query().Get("timeout"); timeoutParam != "" {
		if t, err := strconv.Atoi(timeoutParam); err == nil && t > 0 {
			timeout = time.Duration(t) * time.Second
		}
	}

	message, err := h.service.GetMessage(queueName, timeout)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": message.Data})
}