package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func getCalendarService() *calendar.EventsService {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithScopes(calendar.CalendarScope))
	if err != nil {
		log.Fatal(err)
	}
	return srv.Events
}

func getCalendarEvents(service *calendar.EventsService, calendarId string, eventsType string, initialDate time.Time, weekCount int) []currentEvent {
	out := make([]currentEvent, 0)

	for i := 1; i < weekCount + 1; i++ {
		lowerBound := initialDate.AddDate(0, 0, 7 * (i - 1)).Format(time.RFC3339)
		upperBound := initialDate.AddDate(0, 0, 7 * i).Format(time.RFC3339)
		out = getEventsInPeriod(service, calendarId, lowerBound, upperBound, eventsType, i, out)
	}

	return out
}

func getEventsInPeriod(service *calendar.EventsService, calendarId string, timeMin string, timeMax string,
		evtType string, weekIndex int, outList []currentEvent) []currentEvent {
	events, err := service.List(calendarId).SingleEvents(true).TimeMin(timeMin).TimeMax(timeMax).
		OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve events from calendar: %v", err)
	}
	
	for _, item := range events.Items {
		var startTime, endTime, day string
		loc, _ := time.LoadLocation("Europe/London")
		const timeFmt = "15:04"

		if item.Start.DateTime != "" {
			t, err := time.ParseInLocation(time.RFC3339, item.Start.DateTime, loc)
			if err != nil {
				log.Fatal("Error parsing startTime: ", err)
			}
			startTime = t.Format(timeFmt)
			day = t.Weekday().String()

			t, err = time.ParseInLocation(time.RFC3339, item.End.DateTime, loc)
			if err != nil {
				log.Fatal("Error parsing endTime: ", err)
			}
			endTime = t.Format(timeFmt)
		} else if item.Start.Date == item.End.Date {
			parsedDate, err := time.Parse(time.DateOnly, item.Start.Date)
			if err != nil {
				log.Fatal("Error parsing starting date: ", err)
			}
			day = parsedDate.Weekday().String()
		} else {
			continue
		}

		day = day[:3]

		evt := currentEvent{
			Title: item.Summary,
			Room: item.Location,
			Start: startTime,
			End: endTime,
			Type: evtType,
			Day: day,
			Week: weekIndex,
		}

		outList = append(outList, evt)
	}

	return outList
}

func getAllEvents(initialDate time.Time, weekCount int) []currentEvent {
	calendarMap := map[string]string {
		"academic" : "d6cbb1bdfa033d8c017b90517e977a65c37f7ae9d480778cdceea14db7ab596d@group.calendar.google.com",
		"drinking" : "c7425a49eeb18eed9441c4484e8c3d20af91594af3e024a43dd9d5d33101d08e@group.calendar.google.com",
		"sober" : "fd7a382189416be1f514c1fd77edbb7e09ec05625de041fe4f245593fc9b6d20@group.calendar.google.com",
	}
	out := []currentEvent{}
	service := getCalendarService()

	for eventType, calendarId := range calendarMap {
		events := getCalendarEvents(service, calendarId, eventType, initialDate, weekCount)
		out = append(out, events...)
	}

	return out
}