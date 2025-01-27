package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// WriteCSV creates a CSV file from the provided channels slice
func WriteCSV(filename string, channels []Channel) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		LogMessage(errorLevel, "Failed to create CSV file: "+err.Error())
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"Channel Name", "Team Name", "Channel Type", "Last Update Date", "Last Post Date", "Total Message Count", "Total Message Count Root", "Has Header", "Has Purpose"}
	if err := writer.Write(headers); err != nil {
		LogMessage(errorLevel, "Failed to write CSV headers: "+err.Error())
		return err
	}

	// Write channel data
	for _, channel := range channels {
		record := []string{
			channel.ChannelName,
			channel.TeamName,
			channel.ChannelType,
			channel.LastUpdateDate,
			channel.LastPostDate,
			fmt.Sprintf("%d", channel.TotalMessageCount),
			fmt.Sprintf("%d", channel.TotalMessageCountRoot),
			fmt.Sprintf("%t", channel.HasHeader),
			fmt.Sprintf("%t", channel.HasPurpose),
		}
		if err := writer.Write(record); err != nil {
			errorMsg := fmt.Sprintf("Failed to write CSV record for channel: %s - Error: %s", channel.ChannelName, err.Error())
			LogMessage(errorLevel, errorMsg)
			return err
		}
	}

	return nil
}

// WriteJSON creates a JSON file from the provided channels slice
func WriteJSON(filename string, channels []Channel) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		LogMessage(errorLevel, "Failed to create JSON file: "+err.Error())
		return err
	}
	defer file.Close()

	// Encode channels to JSON and write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON with indentation
	if err := encoder.Encode(channels); err != nil {
		LogMessage(errorLevel, "Failed to write JSON data: "+err.Error())
		return err
	}

	return nil
}
