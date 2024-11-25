package handler

import (
	"be-golang-chapter-36-implem/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	CustomerHandler CustomerHadler
}

func NewAllHandler(service service.AllService, log *zap.Logger) AllHandler {
	return AllHandler{
		CustomerHandler: NewCustomerHandler(service, log),
	}
}
