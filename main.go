package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mattermost/mattermost/server/public/model"
)

var Version = "development" // Default value - overwritten during bild process

type Channel struct {
	ChannelName           string
	ChannelID             string
	TeamID                string
	TeamName              string
	ChannelType           string
	LastUpdateDate        string
	LastPostDate          string
	TotalMessageCount     int64
	TotalMessageCountRoot int64
	HasHeader             bool
	HasPurpose            bool
}

func main() {
	// Parse flags and configuration
	config := parseConfig()

	mmPort := strconv.Itoa(config.Port)

	var dbgCSVString string
	if config.CSV {
		dbgCSVString = "True"
	} else {
		dbgCSVString = "False"
	}

	DebugMessage := fmt.Sprintf("Parameters: \n  MattermostURL=%s\n  MattermostPort=%s\n  MattermostScheme=%s\n  MattermostToken=%s\n  CSV Output=%s\n  Output File=%s\n",
		config.URL,
		mmPort,
		config.Scheme,
		config.Token,
		dbgCSVString,
		config.File,
	)
	DebugPrint(DebugMessage)

	mmTarget := fmt.Sprintf("%s://%s:%s", config.Scheme, config.URL, mmPort)

	DebugPrint("Full target for Mattermost: " + mmTarget)
	mmClient := model.NewAPIv4Client(mmTarget)
	mmClient.SetToken(config.Token)
	DebugPrint("Connected to Mattermost")

	LogMessage(infoLevel, "Processing started - Version: "+Version)

	var channels []Channel

	err := GetChannelStats(*mmClient, &channels)

	if err != nil {
		LogMessage(errorLevel, "Failed to retrieve channel data.  Aborting!")
		os.Exit(10)
	}

	if config.CSV {
		err := WriteCSV(config.File, channels)
		if err != nil {
			LogMessage(errorLevel, "Failed to write CSV file.  Aborting.")
			os.Exit(11)
		}
	} else {
		err := WriteJSON(config.File, channels)
		if err != nil {
			LogMessage(errorLevel, "Failed to write JSON file.  Aborting.")
			os.Exit(12)
		}
	}

	LogMessage(infoLevel, "Channel stats written to: "+config.File)

	os.Exit(0)

}
