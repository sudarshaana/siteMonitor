package main

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/sudarshaana/siteMonitor/cmd/monitor"
	"github.com/sudarshaana/siteMonitor/cmd/web"
	"github.com/sudarshaana/siteMonitor/config"
)

func main() {

	// Load .env file
	config.Load()
	conf := config.GetConfig()

	// slack API init
	api := slack.New(conf.SLACK_API_TOKEN)

	// get the server list array
	servers := monitor.GetServerLists()

	// running server Check on each server
	for _, server := range servers {
		go server.Check(api)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		web.GenerateHTMLResponse(w, servers)
	})

	port := conf.SERVER_PORT
	add := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", add)

	err := http.ListenAndServe(add, mux)
	if err != nil {
		panic(err)

	}

}
