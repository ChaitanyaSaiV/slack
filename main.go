// pkg/slack/main.go
package slack

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
	"time"
)

type Slack_Configs struct {
	SlackConfig *slack.Client
}

func (config *Slack_Configs) Initialize() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	config.SlackConfig = slack.New(os.Getenv("SLACK_TOKEN"))
}

func (config *Slack_Configs) SendWeather(city string, temp string) {
	attachment := slack.Attachment{
		Pretext: "Weather",
		Text:    fmt.Sprintf("Weather in your City : %s is %s", city, temp),
		Color:   "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
		ImageURL: "https://files.slack.com/files-tmb/T06CHSE7Z5L-F06SSAELCP2-b6649a4f5c/image_480.png",
	}
	_, timestamp, err := config.SlackConfig.PostMessage(os.Getenv("SLACK_CHANNEL"), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timestamp)
}
