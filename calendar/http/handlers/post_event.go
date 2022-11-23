package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
	"go.uber.org/zap"
)

type createEventImpl struct {
	logger *zap.SugaredLogger
}

func NewCreateEventHandler(logger *zap.SugaredLogger) operations.PostCreateEventHandler {
	return &createEventImpl{logger}
}

func (c createEventImpl) Handle(params operations.PostCreateEventParams) middleware.Responder {
	c.logger.Infof("request")
	return operations.NewPostCreateEventCreated()
}
