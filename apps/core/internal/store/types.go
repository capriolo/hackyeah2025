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
	PushSubscriptions []PushSubscription `json:"pushSubscriptions"`
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
