package youtube

import (
	"context"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Youtube struct {
	Channels       []string `toml:"channels"`
	APIKey         string   `toml:"api_key"`
	youtubeService *youtube.Service
}

const YoutubeConfig = `
  ## List of channels to monitor
  channels = ["UCqVDpXKLmKeBU_yyt_QkItQ"]
  ## Google API Key
  # api_key = ""
`

func (s *Youtube) SampleConfig() string {
	return YoutubeConfig
}

func (s *Youtube) Description() string {
	return "Gather youtube channel information from Youtube"
}

func (s *Youtube) createYoutubeService(ctx context.Context) (*youtube.Service, error) {
	return youtube.NewService(ctx, option.WithAPIKey(s.APIKey))
}

func (s *Youtube) Gathertrig(acc telegraf.Accumulator) error {
	ctx := context.Background()

	if s.youtubeService == nil {
		service, err := s.createYoutubeService(ctx)
		if err != nil {
			return err
		}
		s.youtubeService = service
	}

	strArray := []string{"snippet", "statistics"}
	call := s.youtubeService.Channels.List(strings.Join(strArray, ",")).Id(strings.Join(s.Channels, ",")).MaxResults(50)

	resp, err := call.Do()

	if err != nil {
		return err
	}

	now := time.Now()

	for _, item := range resp.Items {
		tags := getTags(item)
		fields := getFields(item)

		acc.AddFields("youtube_channel", fields, tags, now)
	}

	return nil
}

func getTags(channelInfo *youtube.Channel) map[string]string {
	return map[string]string{
		"id":    channelInfo.Id,
		"title": channelInfo.Snippet.Title,
	}
}

func getFields(channelInfo *youtube.Channel) map[string]interface{} {
	return map[string]interface{}{
		"subscribers": channelInfo.Statistics.SubscriberCount,
		"videos":      channelInfo.Statistics.VideoCount,
		"views":       channelInfo.Statistics.ViewCount,
		"country":     channelInfo.Snippet.Country,
		"channelName": channelInfo.Snippet.Title,
		"createdon":   channelInfo.Snippet.PublishedAt,
	}
}

func init() {
	inputs.Add("youtube", func() telegraf.Input {
		return &Youtube{}
	})
}
