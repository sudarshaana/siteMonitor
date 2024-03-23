package monitor

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/nlopes/slack"
	"github.com/sudarshaana/siteMonitor/config"
)

const DefaultCheckPeriod = 5 * time.Second

type Server struct {
	URL            string
	Name           string
	SlackChannelID string
	Status         string
	LastCheck      time.Time
	CheckPeriod    time.Duration
}

func (s *Server) Check(api *slack.Client) {
	if s.CheckPeriod <= 0 {
		s.CheckPeriod = DefaultCheckPeriod
	}

	ticker := time.NewTicker(s.CheckPeriod)
	defer ticker.Stop()

	for {
		ok, statusCode := s.doCheck()
		if !ok {
			//s.Status = fmt.Sprintf("❌ Not Responding (StatusCode: %d)", statusCode)
			// send slack notification
			s.sendSlackNotification(api)
		} else {
			s.LastCheck = time.Now()
			//s.Status = fmt.Sprintf("✅%d", statusCode)
			// fmt.Printf("%s --- %s\n", s.Name, s.CheckPeriod.String())
		}
		s.Status = fmt.Sprintf("%d", statusCode)
		<-ticker.C
	}
}

func (s *Server) doCheck() (bool, int) {
	timeout := config.GetConfig().REQUEST_TIMEOUT
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, err := client.Get(s.URL)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// check for google, if ping is possible.
			monitorServiceOk := pingGoogleSuccessfully()
			if !monitorServiceOk {
				// error in monitor service, Handle separately.
				return false, 0

			} else {
				fmt.Printf("timeout error while checking server %s\n", s.Name)
				return false, -1
			}
		}
		fmt.Printf("Error while checking server %s: %s\n", s.Name, err)
		return false, 0
	}

	defer resp.Body.Close()
	responseCode := resp.StatusCode
	return responseCode >= 200 && responseCode < 300, responseCode
}

func (s *Server) sendSlackNotification(api *slack.Client) {
	conf := config.GetConfig()
	// if turned off, won't send msg but print a log
	if !conf.SEND_SLACK_NOTIFICATION {
		fmt.Println("Sending message to Slack is turned off in the .env file.")
		return
	}
	message := fmt.Sprintf("❌ `%s` not responding properly! StatusCode: `%s`", s.Name, s.Status)
	_, _, err := api.PostMessage(s.SlackChannelID, slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("Error sending message Slack: %s\n", err)
	}
}

func pingGoogleSuccessfully() bool {
	// to be sure that problem is not in our end.
	_, err := http.Get("https://google.com")
	return err == nil

}

func GetServerLists() []*Server {
	conf := config.GetConfig()
	// to get notification in different slack channel change the `SlackChannelID`
	servers := []*Server{
		{
			URL:            "https://google.com/",
			Name:           "Google",
			SlackChannelID: conf.SLACK_CHANNEL_ID,
			CheckPeriod:    3 * time.Second,
		},
		{
			URL:            "https://facebook.com/",
			Name:           "Facebook",
			SlackChannelID: conf.SLACK_CHANNEL_ID,
			CheckPeriod:    6 * time.Second,
		},
		{
			URL:            "http://https://invaliddomain.com/",
			Name:           "invalid domain",
			SlackChannelID: conf.SLACK_CHANNEL_ID,
			CheckPeriod:    7 * time.Second,
		},
	}
	return servers
}
