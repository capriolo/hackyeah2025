package sync

import (
	"fmt"
	"net/http"
	"os"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/capriolo/hackyeah2025/core/internal/store"
)

func SyncIcal(user string) {

	data := store.Read()
	profile := data.UserProfile(user)

	profile.CalendarUrl = os.Getenv("ICAL_URL")
	data.SaveUserProfile(profile)

	var events []store.Event

	if profile.CalendarUrl != "" {
		resp, err := http.Get(profile.CalendarUrl)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		cal, err := ics.ParseCalendar(resp.Body)
		if err != nil {
			panic(err)
		}
		// Iteracja po wydarzeniach
		for _, e := range cal.Events() {
			events = append(events, ConvertIcalEvent(e))
		}
	}

	data.SaveUserCalendarEvents(user, events)
	store.Write(data)
}

func ConvertIcalEvent(e *ics.VEvent) store.Event {
	// {
	// 			Type:        "gym",
	// 			Title:       "Siłownia",
	// 			Description: "Trening klatki i tricepsów",
	// 			Start:       330,
	// 			Duration:    100,
	// 		},
	// 		{
	// 			Type:        "work",
	// 			Title:       "Programowanie",
	// 			Description: "Praca nad frontendem aplikacji na hackYeah",
	// 			Start:       480,
	// 			Duration:    100,
	// 		},
	// 		{
	// 			Type:     "travel",
	// 			Title:    "Droga do pracy",
	// 			Start:    420,
	// 			Duration: 50,
	// 		},
	// 		{
	// 			Type:        "meeting",
	// 			Title:       "Spotkanie z Michałem",
	// 			Description: "Plan nowego projektu. Pamiętaj aby zabrać laptopa",
	// 			Start:       700,
	// 			Duration:    240,
	// 		},

	start, _ := timestampToMinutesFromMidnight(e.GetProperty(ics.ComponentPropertyDtStart).Value)
	duration, _ := timestampToMinutesFromMidnight(e.GetProperty(ics.ComponentPropertyDtEnd).Value)
	eDate, _ := timestampToDate(e.GetProperty(ics.ComponentPropertyDtStart).Value)

	return store.Event{
		Id:          e.GetProperty(ics.ComponentPropertyUniqueId).Value,
		Date:        eDate,
		Type:        "sleepTime",
		Title:       e.GetProperty(ics.ComponentPropertySummary).Value,
		Description: "Jutro masz trening z samego rana. Zapenij wystarczającą ilość snu",
		Start:       start,
		Duration:    duration - start,
	}
}

func timestampToDate(timestamp string) (date string, err error) {
	t, err := time.Parse("20060102T150405Z", timestamp)

	if err != nil {
		return "", fmt.Errorf("invalid timestamp format: %v", err)
	}

	return t.Format("20060102"), nil
}

func timestampToMinutesFromMidnight(timestamp string) (minutesFromMidnight uint16, err error) {
	t, err := time.Parse("20060102T150405Z", timestamp)
	if err != nil {
		return 0, fmt.Errorf("invalid timestamp format: %v", err)
	}

	minutesFromMidnight = uint16(t.Hour()*60 + t.Minute() + 120)

	return minutesFromMidnight, nil
}
