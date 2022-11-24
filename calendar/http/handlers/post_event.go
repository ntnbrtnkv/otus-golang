package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
	"github.com/ntnbrtnkv/otus-golang/calendar/model"
	"go.uber.org/zap"
	"time"
)

type createEventImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewCreateEventHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.PostCreateEventHandler {
	return &createEventImpl{logger, calendar}
}

func (c createEventImpl) Handle(params operations.PostCreateEventParams) middleware.Responder {
	c.logger.Infof("Create event")

	uuidGen, err := uuid.NewUUID()

	if err != nil {
		c.logger.Infof("Failed to create event: %s", err)
		return operations.NewPostCreateEventBadRequest().WithPayload(&models.Error{
			Error: fmt.Sprint(err),
		})
	}

	event := model.Event{
		ID:          uuidGen.String(),
		TimeStart:   time.Time(params.TimeStart),
		TimeEnd:     time.Time(params.TimeEnd),
		Title:       params.Title,
		Description: params.Description,
	}

	c.logger.Debugf("%v", event)

	if err = c.calendar.AddEvent(event); err != nil {
		c.logger.Infof("Failed to create event: %s", err)
		return operations.NewPostCreateEventBadRequest().WithPayload(&models.Error{
			Error: fmt.Sprint(err),
		})
	}

	return operations.NewPostCreateEventCreated()
}
