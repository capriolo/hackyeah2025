package day

import (
	"fmt"
	"time"

	"github.com/capriolo/hackyeah2025/core/internal/store"
)

type DayRes struct {
	Date        string          `json:"date"`
	Description string          `json:"description"`
	Events      []*store.Event  `json:"events"`
	Suggestions []*store.Event  `json:"suggestions"`
	WakeupTime  uint16          `json:"wakeupTime,omitempty"`
	SleepTime   uint16          `json:"sleepTime,omitempty"`
	SleepData   store.SleepData `json:"sleepData,omitempty"`
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

	sleepData := &store.SleepData{}
	for _, sd := range user.SleepData {
		if sd.Date == dateStr {
			sleepData = sd
			break
		}
	}

	exampleData := exampleData(dateStr)

	return DayRes{
		Description: exampleData.Description,
		Events:      events,
		Suggestions: exampleData.Suggestions,
		WakeupTime:  exampleData.WakeupTime,
		SleepTime:   exampleData.SleepTime,
		SleepData:   *sleepData,
	}
}

func exampleData(date string) DayRes {
	switch date {
	case "20251005":
		return DayRes{
			WakeupTime:  0,
			SleepTime:   0,
			Description: "Po zarwanej nocy postaraj się wcześniej położyć aby odespać. Aby zniwelować skutki zmęczenia po HackYeah możesz zastosować mega dawkę kreatyny (20g).",
			Suggestions: []*store.Event{
				{
					Type:        "wakeUp",
					Title:       "Czas wstawać",
					Description: "Mam nadzieję, że udało sie chociaż trochę nadrobić zarwaną noc.",
					Start:       480,
				},
				{
					Type:        "pills",
					Title:       "Suplementacja kreatyną",
					Description: "Duża dawka może fajnie zadziałać na sprawność umysłu. Zastosuj zamiast kawy.",
					Start:       525,
				},
				{
					Type:        "food",
					Title:       "Lunch",
					Description: "Zaplanuj pełnowartościowy posiłek",
					Start:       720,
				},
				{
					Type:        "glasses",
					Title:       "Okulary AntiBlue",
					Description: "Zredukuj ilość światła niebieskiego aby poprawić jakość snu",
					Start:       1100,
				},
				{
					Type:        "sleepTime",
					Title:       "Pora spać",
					Description: "Dobrej nocy. Pamiętaj aby zaciemnić sypialnie",
					Start:       1200,
				},
			},
		}
	case "20251006":
		return DayRes{
			WakeupTime:  0,
			SleepTime:   0,
			Description: "Tak długi trening po ciężkim weekendzie nie jest dobrym pomysłem. Sugeruję go przełożyć na następny dzień. Podczas pracy biurowej znajdź chwilę na sycący posiłek",
			Suggestions: []*store.Event{
				{
					Type:        "wakeUp",
					Title:       "Pobudka",
					Description: "Pora zbierać się na trening.",
					Start:       300,
				},
				{
					Type:        "food",
					Title:       "Posiłek potreningowy",
					Description: "Po ciężkim treningu siłowym i biegowym uzupełnij glikogen w mięśniach. Nie bój się porcji węglowodanów nawet jeśli jesteś w trybie Low Carb",
					Start:       490,
				},
				{
					Type:        "food",
					Title:       "Lunch",
					Description: "Zaplanuj pełnowartościowy posiłek",
					Start:       780,
				},
				{
					Type:        "food",
					Title:       "Kolacja",
					Description: "Zaplanuj pełnowartościowy posiłek",
					Start:       1050,
				},
				{
					Type:        "glasses",
					Title:       "Okulary AntiBlue",
					Description: "Zredukuj ilość światła niebieskiego aby poprawić jakość snu",
					Start:       1200,
				},
				{
					Type:        "sleepTime",
					Title:       "Pora spać",
					Description: "Dobrej nocy. Pamiętaj aby zaciemnić sypialnie",
					Start:       1310,
				},
			},
		}
	case "20251007":
		return DayRes{
			WakeupTime:  0,
			SleepTime:   0,
			Description: "Dziś w planie dłuższa podróż, zabierz ze sobą jakąś przekąskę. Kabanosy lub batony proteinowe mogą być dobrym pomysłem aby doładować białko po mocnych treningach.",
			Suggestions: []*store.Event{
				{
					Type:        "wakeUp",
					Title:       "Pobudka",
					Description: "Pora zbierać się na trening.",
					Start:       340,
				},
				{
					Type:        "food",
					Title:       "Posiłek potreningowy",
					Description: "Po treningu siłowym zjedz coś dobrego",
					Start:       450,
				},
				{
					Type:        "coffee",
					Title:       "Kawa",
					Description: "Spotkanie biznesowe to dobra pora na kawę.",
					Start:       510,
				},
				{
					Type:        "food",
					Title:       "Lunch",
					Description: "Zaplanuj pełnowartościowy posiłek",
					Start:       720,
				},
				{
					Type:        "glasses",
					Title:       "Okulary AntiBlue",
					Description: "Zredukuj ilość światła niebieskiego aby poprawić jakość snu",
					Start:       1150,
				},
				{
					Type:        "sleepTime",
					Title:       "Pora spać",
					Description: "Po długiej podróży zadbaj o regeneracyjny sen",
					Start:       1260,
				},
			},
		}
	default:
		return DayRes{
			WakeupTime:  0,
			SleepTime:   0,
			Description: "Sen nie był najlepszy, a kalendarz wygląda napięcie. Rozważ krótki 20-minutowy power nap po lunchu — poprawi koncentrację bardziej niż kolejna kawa.",
			Suggestions: []*store.Event{
				{
					Type:        "wakeUp",
					Title:       "Pobudka",
					Description: "Pora zbierać się na trening.",
					Start:       340,
				},
				{
					Type:        "food",
					Title:       "Posiłek",
					Description: "Zjedz posiłek 800kcal, 50g białka",
					Start:       500,
					Duration:    10,
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
		}
	}
}
