package day

type Event struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Start       uint16 `json:"start"`
	Duration    uint16 `json:"duration"`
}

type DayRes struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Events      []Event `json:"events"`
	Suggestions []Event `json:"suggestions"`
	WakeupTime  uint16  `json:"wakeupTime,omitempty"`
	SleepTime   uint16  `json:"sleepTime,omitempty"`
}

func Day(date string) DayRes {
	return DayRes{
		Description: "Sen nie był najlepszy, a kalendarz wygląda napięcie. Rozważ krótki 20-minutowy power nap po lunchu — poprawi koncentrację bardziej niż kolejna kawa.",
		Events: []Event{
			{
				Type:        "gym",
				Title:       "Siłownia",
				Description: "Trening klatki i tricepsów",
				Start:       330,
				Duration:    100,
			},
			{
				Type:        "work",
				Title:       "Programowanie",
				Description: "Praca nad frontendem aplikacji na hackYeah",
				Start:       480,
				Duration:    100,
			},
			{
				Type:     "travel",
				Title:    "Droga do pracy",
				Start:    420,
				Duration: 50,
			},
			{
				Type:        "meeting",
				Title:       "Spotkanie z Michałem",
				Description: "Plan nowego projektu. Pamiętaj aby zabrać laptopa",
				Start:       700,
				Duration:    240,
			},
		},
		Suggestions: []Event{
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
