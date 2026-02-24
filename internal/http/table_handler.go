package http

import (
	"encoding/json"
	"net/http"

	"github.com/arizaguaca/table/internal/domain"
)

type TableHandler struct {
	Usecase domain.TableUsecase
}

func NewTableHandler(u domain.TableUsecase) *TableHandler {
	return &TableHandler{
		Usecase: u,
	}
}

func (h *TableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var table domain.Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Usecase.Create(r.Context(), &table); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(table)
}

func (h *TableHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	tables, err := h.Usecase.Fetch(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}
