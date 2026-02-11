package api

import (
	"github.com/sssseraphim/microservices/order/internal/application/core/domain"
	"github.com/sssseraphim/microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db.
