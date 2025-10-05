package day

import (
	"fmt"
	"time"

	"github.com/capriolo/hackyeah2025/core/internal/store"
)

type DayRes struct {
	Date        string         `json:"date"`
	Description string         `json:"description"`
	Events      []*store.Event `json:"events"`
	Suggestions []*store.Event `json:"suggestions"`
	WakeupTime  uint16         `json:"wakeupTime,omitempty"`
	SleepTime   uint16         `json:"sleepTime,omitempty"`
}

func Day(user *store.User, dateStr string) DayRes {

	today := time.Now()

	if dateStr == "" {
		dateStr = today.Format("20060102")
	}

	date, err := time.Parse("20060102", dateStr)

	if date.Before(today) || date.Equal(today) {
		fmt.Print("grab sleep data")
	}

	if err != nil {
		panic(err)
	}

	var events []*store.Event

	for _, e := range user.CalendarEvents {
		if dateStr == e.Date {
			events = append(events, e)
		}
	}

	return DayRes{
		Description: "Sen nie był najlepszy, a kalendarz wygląda napięcie. Rozważ krótki 20-minutowy power nap po lunchu — poprawi koncentrację bardziej niż kolejna kawa.",
		Events:      events,
		Suggestions: []*store.Event{
			{
				Type:        "coffee",
				Title:       "Kawa",
				Description: "Kawa przed spotkaniem pozwoli nabrać energii",
				Start:       600,
				Duration:    10,
			},
			{
				Type:        "food",
				Title:       "Posiłek",
				Description: "Zjedz posiłek 800kcal, 50g białka",
				Start:       500,
				Duration:    10,
			},
			{
				Type:        "pills",
				Title:       "Suplementacja magnezu",
				Description: "Magnez na noc wchłonie się w większej ilości",
				Start:       1000,
			},
			{
				Type:        "glasses",
				Title:       "Zredukuj niebieskie światło",
				Description: "Jeśli musisz korzystać z ekranów załóż okulary z filtrem światła niebieskiego",
				Start:       860,
			},
			{
				Type:        "sleepTime",
				Title:       "Pora spać",
				Description: "Jutro masz trening z samego rana. Zapenij wystarczającą ilość snu",
				Start:       1020,
			},
		},
		WakeupTime: 300,
		SleepTime:  1200,
	}
}
