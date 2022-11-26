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

type getEventsForMonthImpl struct {
	logger   *zap.SugaredLogger
	calendar model.Calendar
}

func NewGetEventsForMonthHandler(logger *zap.SugaredLogger, calendar model.Calendar) operations.GetEventsForMonthHandler {
	return &getEventsForMonthImpl{logger, calendar}
}

func (c getEventsForMonthImpl) Handle(params operations.GetEventsForMonthParams) middleware.Responder {
	year := time.Time(params.Month).Year()
	month := time.Time(params.Month).Month()

	c.logger.Infof("Get events for month: %d-%s", year, month)

	events := c.calendar.ListEvents()
	var response = make([]*models.Event, 0)

	for _, event := range events {
		tsYear := event.TimeStart.Year()
		tsMonth := event.TimeStart.Month()
		teYear := event.TimeEnd.Year()
		teMonth := event.TimeEnd.Month()
		if (tsYear < year || tsYear == year && tsMonth <= month) && (teYear > year || teYear == year && teMonth >= month) {
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
