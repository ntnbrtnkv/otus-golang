package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
	"github.com/ntnbrtnkv/otus-golang/calendar/model"
	"go.uber.org/zap"
	"time"
)

type getEventsForDayImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewGetEventsForDayHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.GetEventsForDayHandler {
	return &getEventsForDayImpl{logger, calendar}
}

func (c getEventsForDayImpl) Handle(params operations.GetEventsForDayParams) middleware.Responder {
	c.logger.Infof("Get events for day: %s", params.Day)

	year := time.Time(params.Day).Year()
	day := time.Time(params.Day).YearDay()

	events := c.calendar.ListEvents()
	var response = make([]*models.Event, 0)

	for _, event := range events {
		tsYear := event.TimeStart.Year()
		tsDay := event.TimeStart.YearDay()
		teYear := event.TimeEnd.Year()
		teDay := event.TimeEnd.YearDay()
		if (tsYear < year || tsYear == year && tsDay <= day) && (teYear > year || teYear == year && teDay >= day) {
			ts := strfmt.DateTime(event.TimeStart)
			te := strfmt.DateTime(event.TimeEnd)
			response = append(response, &models.Event{
				ID:          event.ID,
				TimeStart:   &ts,
				TimeEnd:     &te,
				Title:       &event.Title,
				Description: &event.Description,
			})
		}
	}

	return operations.NewGetEventsForDayOK().WithPayload(&models.Events{Result: response})
}
