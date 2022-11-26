package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
	"github.com/ntnbrtnkv/otus-golang/calendar/model"
	"go.uber.org/zap"
)

type deleteEventImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewDeleteEventHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.PostDeleteEventHandler {
	return &deleteEventImpl{logger, calendar}
}

func (c deleteEventImpl) Handle(params operations.PostDeleteEventParams) middleware.Responder {
	c.logger.Infof("Delete event: %s", params.ID)

	err := c.calendar.RemoveEvent(params.ID)

	if err != nil {
		c.logger.Infof("Failed to delete event: %s", err)
		return operations.NewPostDeleteEventBadRequest().WithPayload(&models.Error{
			Error: fmt.Sprint(err),
		})
	}

	return operations.NewPostDeleteEventOK()
}
