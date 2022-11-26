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

type getEventsForWeekImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewGetEventsForWeekHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.GetEventsForWeekHandler {
	return &getEventsForWeekImpl{logger, calendar}
}

func (c getEventsForWeekImpl) Handle(params operations.GetEventsForWeekParams) middleware.Responder {
	year, week := time.Time(params.Week).ISOWeek()

	c.logger.Infof("Get events for week: %d-%d", year, week)

	events := c.calendar.ListEvents()
	var response = make([]*models.Event, 0)

	for _, event := range events {
		tsYear, tsWeek := event.TimeStart.ISOWeek()
		teYear, teWeek := event.TimeEnd.ISOWeek()
		if (tsYear < year || tsYear == year && tsWeek <= week) && (teYear > year || teYear == year && teWeek >= week) {
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
