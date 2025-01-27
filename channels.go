package main

import (
	"context"
	"errors"
	"time"

	"github.com/mattermost/mattermost/server/public/model"
)

const (
	pageSize = 50
)

// ConvertEpochToString converts a millisecond epoch time to a formatted string
func ConvertEpochToString(epochMillis int64) string {
	// Convert milliseconds to seconds and nanoseconds
	seconds := epochMillis / 1000
	nanoseconds := (epochMillis % 1000) * int64(time.Millisecond)

	// Create a time.Time object
	t := time.Unix(seconds, nanoseconds)

	// Format the time to a standard string (e.g., RFC3339 format)
	return t.Format(time.RFC3339)
}

func StringContainsData(str string) bool {
	return len(str) > 0
}

func GetChannelStats(mmClient model.Client4, channels *[]Channel) error {
	DebugPrint("Getting channel stats")

	ctx := context.Background()
	etag := ""

	page := 0

	for {
		allchannels, response, err := mmClient.GetAllChannels(ctx, page, pageSize, etag)
		if err != nil {
			LogMessage(errorLevel, "Failed to retrieve channels: "+err.Error())
			return err
		}
		if response.StatusCode != 200 && response.StatusCode != 201 {
			LogMessage(errorLevel, "Function call to GetAllChannels returned bad HTTP response")
			return errors.New("bad HTTP response")
		}

		// Exit the loop if we're not getting any more channels returned
		if len(allchannels) == 0 {
			break
		}

		for _, channel := range allchannels {
			*channels = append(*channels, Channel{
				ChannelName:           channel.DisplayName,
				ChannelID:             channel.Id,
				TeamID:                channel.TeamId,
				TeamName:              channel.TeamDisplayName,
				ChannelType:           string(channel.Type),
				LastUpdateDate:        ConvertEpochToString(channel.UpdateAt),
				LastPostDate:          ConvertEpochToString(channel.LastPostAt),
				TotalMessageCount:     channel.TotalMsgCount,
				TotalMessageCountRoot: channel.TotalMsgCountRoot,
				HasHeader:             StringContainsData(channel.Header),
				HasPurpose:            StringContainsData(channel.Purpose),
			})
		}

		page++
	}

	return nil
}
