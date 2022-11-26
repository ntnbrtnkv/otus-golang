package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
	"github.com/ntnbrtnkv/otus-golang/calendar/model"
	"go.uber.org/zap"
	"time"
)

type updateEventImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewUpdateEventHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.PostUpdateEventHandler {
	return &updateEventImpl{logger, calendar}
}

func (c updateEventImpl) Handle(params operations.PostUpdateEventParams) middleware.Responder {
	c.logger.Infof("Update event: %s", params.ID)

	event := model.Event{
		ID:          params.ID,
		TimeStart:   time.Time(params.TimeStart),
		TimeEnd:     time.Time(params.TimeEnd),
		Title:       params.Title,
		Description: params.Description,
	}

	err := c.calendar.ChangeEvent(event)

	if err != nil {
		c.logger.Infof("Failed to update event: %s", err)
		return operations.NewPostUpdateEventBadRequest().WithPayload(&models.Error{
			Error: fmt.Sprint(err),
		})
	}

	c.logger.Debugf("Updated event %v", event)

	return operations.NewPostUpdateEventOK()
}
