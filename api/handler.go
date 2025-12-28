package api

import (
	"encoding/json"
	"net/http"
	"orchestrator_llm/domain"
	"orchestrator_llm/orchestrator"
)

type Handler struct {
	Service *orchestrator.Service
}

func (h *Handler) HealthOK(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}

func (h *Handler) Ask(w http.ResponseWriter, r *http.Request) {
	var req domain.AskRequest
	json.NewDecoder(r.Body).Decode(&req)

	answer, err := h.Service.Ask(req)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(domain.AskResponse{
		Answer: answer,
	})

}
