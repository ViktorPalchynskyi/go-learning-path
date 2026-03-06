package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_13/internal/domain"
	"github.com/username/lesson_13/internal/service"
)

type CreateAccountRequest struct {
	OwnerId int64 `json:"owner_id"`
	Balance float64 `json:"balance"`
}

type TransferRequest struct {
	FromId int64 `json:"from_id"`
	ToId int64 `json:"to_id"`
	Amount float64 `json:"amount"`
}

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler{
	return &AccountHandler{service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}

	account, err := h.service.Create(ctx, req.OwnerId, req.Balance)
	if err != nil {
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, 201, account)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64);
	if err !=nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid id"})
		return
	}

	account, err := h.service.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeJSON(w, 404, ErrorResponse{Error: "not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, account)
}

func (h *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}

	err := h.service.Transfer(ctx, req.FromId, req.ToId, req.Amount)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeJSON(w, 404, ErrorResponse{Error: "not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, nil)
}