package store

type PushSubscriptionKeys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type PushSubscription struct {
	Endpoint string               `json:"endpoint"`
	Keys     PushSubscriptionKeys `json:"keys"`
}

type User struct {
	Name              string             `json:"name"`
	CalendarUrl       string             `json:"calendarUrl"`
	RunalyzeToken     string             `json:"runalyzeToken"`
	CalendarEvents    []*Event           `json:"calendarEvents"`
	SleepData         []*SleepData       `json:"sleepData"`
	PushSubscriptions []PushSubscription `json:"pushSubscriptions"`
}

type SleepData struct {
	Id                 int    `json:"id"`
	Date               string `json:"date"`
	Duration           uint16 `json:"duration"`
	RemDuration        uint16 `json:"remDuration"`
	LightSleepDuration uint16 `json:"lightSleepDuration"`
	DeepSleepDuration  uint16 `json:"deepSleepDuration"`
	AwakeDuration      uint16 `json:"awakeDuration"`
}

type Event struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Start       uint16 `json:"start"`
	Duration    uint16 `json:"duration"`
}

type Data struct {
	Users []User `json:"users"`
}

func (d *Data) AddUserSubscription(user string, sub PushSubscription) {

	mu.Lock()
	defer mu.Unlock()

	for i, u := range d.Users {
		if u.Name == user {
			d.Users[i].PushSubscriptions = append(d.Users[i].PushSubscriptions, sub)
			return
		}
	}

	d.Users = append(d.Users, User{
		Name:              user,
		PushSubscriptions: []PushSubscription{sub},
	})
}

func (d *Data) UserSubscriptions(user string) []PushSubscription {

	mu.RLock()
	defer mu.RUnlock()

	for i, u := range d.Users {
		if u.Name == user {
			return d.Users[i].PushSubscriptions
		}
	}

	return []PushSubscription{}
}

func (d *Data) UserProfile(user string) User {

	mu.RLock()
	defer mu.RUnlock()

	for i, u := range d.Users {
		if u.Name == user {
			return d.Users[i]
		}
	}

	return User{}
}

func (d *Data) SaveUserSleepData(user string, sleepData []*SleepData) {
	mu.RLock()
	defer mu.RUnlock()

	for i, u := range d.Users {
		if u.Name == user {
		NewSleepData:
			for _, ne := range sleepData {

				for _, ee := range d.Users[i].SleepData {
					if ee.Id == ne.Id {
						continue NewSleepData
					}
				}

				d.Users[i].SleepData = append(d.Users[i].SleepData, ne)
			}

			break
		}
	}
}

func (d *Data) SaveUserCalendarEvents(user string, events []Event) {
	mu.RLock()
	defer mu.RUnlock()

	for i, u := range d.Users {
		if u.Name == user {
		NewEvents:
			for _, ne := range events {

				for _, ee := range d.Users[i].CalendarEvents {
					if ee.Id == ne.Id {
						continue NewEvents
					}
				}

				d.Users[i].CalendarEvents = append(d.Users[i].CalendarEvents, &ne)
			}

			break
		}
	}
}

func (d *Data) SaveUserProfile(user User) {

	mu.RLock()
	defer mu.RUnlock()

	for i, u := range d.Users {
		if u.Name == user.Name {
			d.Users[i].CalendarUrl = user.CalendarUrl

			return
		}
	}

	d.Users = append(d.Users, user)
}
