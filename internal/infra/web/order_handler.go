package web

import (
	"encoding/json"
	"net/http"

	"github.com/backendengineerark/clean-arch/internal/entity"
	"github.com/backendengineerark/clean-arch/internal/usecase"
	"github.com/backendengineerark/clean-arch/pkg/events"
)

type WebOrderHandler struct {
	OrderRepository   entity.OrderRepositoryInterface
	EventDispatcher   events.EventDispatcherInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	orderRepository entity.OrderRepositoryInterface,
	eventDispatcher events.EventDispatcherInterface,
	orderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository:   orderRepository,
		EventDispatcher:   eventDispatcher,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrderUseCase := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrderUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
