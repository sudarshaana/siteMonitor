package web

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/sudarshaana/siteMonitor/cmd/monitor"
)

func GenerateHTMLResponse(w http.ResponseWriter, servers []*monitor.Server) {

	indexFileLocation := filepath.Join("static", "index.html")
	file, err := os.Open(indexFileLocation)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"GetIcon":         GetIcon,
		"convertDateTime": convertDateTime,
	}).ParseFiles(indexFileLocation))

	if err := tmpl.Execute(w, servers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetIcon(status string) string {
	// Determine the icon based on the status value
	if strings.HasPrefix(status, "0") {
		return "ℹ️"
	} else if strings.HasPrefix(status, "2") {
		return "✅"
	} else {
		return "❌"
	}
}

func convertDateTime(timeObj time.Time) string {
	bangladeshLocation, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		fmt.Println("Error loading Bangladesh timezone:", err)
		return "x"
	}
	formattedTimeObj := timeObj.In(bangladeshLocation)
	humanReadableTime := formattedTimeObj.Format("02-January-2006, 03:04:05 PM")

	// diff := time.Now().Sub(timeObj)
	// lastChecked := fmt.Sprintf("%s | %s", humanReadableTime, diff)
	return humanReadableTime
}
