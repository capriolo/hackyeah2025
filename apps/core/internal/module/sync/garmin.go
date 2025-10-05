package sync

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/capriolo/hackyeah2025/core/internal/store"
)

func SyncGarmin(user string) {

	data := store.Read()
	profile := data.UserProfile(user)

	profile.CalendarUrl = os.Getenv("RUNALYZE_TOKEN")
	data.SaveUserProfile(profile)
	sleepData := GetData()
	data.SaveUserSleepData(user, sleepData)

	store.Write(data)
}

func GetData() []*store.SleepData {

	allData := []*store.SleepData{}

	page := uint16(1)
	for {
		sleepData := runalyzeCall(page)

		if len(sleepData) == 0 {
			break
		}

		if page >= 6 {
			break
		}
		limit, _ := time.Parse("20060102", "20251001")

		for _, entry := range sleepData {
			d, _ := time.Parse("20060102", entry.DateTime)

			if d.Before(limit) {
				continue
			}

			allData = append(allData, convertRunalyzeData(entry))
		}

		page++
	}

	return allData
}

func convertRunalyzeData(rd runalyzeSleepData) *store.SleepData {

	t, _ := time.Parse(time.RFC3339, rd.DateTime)

	// Konwersja na format YYYYMMDD
	date := t.Format("20060102")
	return &store.SleepData{
		Id:                 rd.Id,
		Date:               date,
		Duration:           rd.Duration,
		RemDuration:        rd.RemDuration,
		LightSleepDuration: rd.LightSleepDuration,
		DeepSleepDuration:  rd.DeepSleepDuration,
		AwakeDuration:      rd.AwakeDuration,
	}
}

func runalyzeCall(page uint16) []runalyzeSleepData {
	url := fmt.Sprintf("https://runalyze.com/api/v1/metrics/sleep?page=%d", page)

	// Tworzenie requesta
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Ustawienie nagłówków
	req.Header.Set("accept", "application/json")
	req.Header.Set("token", os.Getenv("RUNALYZE_TOKEN"))

	// Wysłanie requesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Sprawdzenie statusu odpowiedzi
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("API returned status %d: %s", resp.StatusCode, string(body))
		return []runalyzeSleepData{}
	}

	// Odczytanie odpowiedzi
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Parsowanie JSON
	var sleepData []runalyzeSleepData
	err = json.Unmarshal(body, &sleepData)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	return sleepData
}

type runalyzeSleepData struct {
	Id                 int    `json:"id"`
	DateTime           string `json:"date_time"`
	Duration           uint16 `json:"duration"`
	RemDuration        uint16 `json:"rem_duration"`
	LightSleepDuration uint16 `json:"light_sleep_duration"`
	DeepSleepDuration  uint16 `json:"deep_sleep_duration"`
	AwakeDuration      uint16 `json:"awake_duration"`
}
