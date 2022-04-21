package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ComputeWeeksLeftBeforeNextYear() (weeksLeft string, err error) {
	response, err := http.Get("https://worldtimeapi.org/api/timezone/Europe/Paris")
	if err != nil || response == nil {
		return "Call failed", err
	}
	defer response.Body.Close()

	if response.Status != "200 OK" {
		return "Response failed", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "Read failed", err
	}

	timeResponse := &TimeResponse{}
	err = json.Unmarshal(body, timeResponse)
	if err != nil {
		return "Parse failed", err
	}

	return fmt.Sprintf("~ %d weeks left this year", 52-timeResponse.WeekNumber), err
}

type TimeResponse struct {
	WeekNumber int `json:"week_number"`
}
