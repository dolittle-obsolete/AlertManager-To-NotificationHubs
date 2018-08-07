package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	configuration "../configuration"
)

// WebHook is the received WebHook message
type WebHook struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}

// Alert
type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
}

var (
	// Config is the configuration for the server
	Config configuration.Config
)

// StartServer starts the server
func StartServer(config configuration.Config) {
	Config = config

	portStr := strconv.FormatInt(config.ServerPort, 10)
	serverAddrStr := config.ServerAddress + ":" + portStr
	http.HandleFunc("/webhook", hookHandler)

	log.Printf("Listening at %s for inncomming webhook", serverAddrStr)
	log.Fatal(http.ListenAndServe(serverAddrStr, nil))
}

func hookHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Inncomming request")

	if request.Method != http.MethodPost {
		http.Error(writer, "Server only accepts POST requests.", http.StatusMethodNotAllowed)
		return
	}
	var webHook WebHook
	request.ParseForm()

	logHeaderIfDebug(request)
	err := json.NewDecoder(request.Body).Decode(&webHook)
	if err != nil {
		badRequest(writer, err)
		return
	}
	logPayloadIfDebug(&webHook)
	logWebHookIfDebug(&webHook)

	writer.WriteHeader(http.StatusOK)
	log.Println("Request handled")
}

func badRequest(writer http.ResponseWriter, err error) {
	http.Error(writer, "Bad request. Error: ` + err.Error()", http.StatusBadRequest)
}

func logHeaderIfDebug(request *http.Request) {
	if Config.Debug {
		log.Println("Request header:")
		log.Printf("%s\n", request.Header)
	}
}

func logPayloadIfDebug(webHook *WebHook) {
	if Config.Debug {
		bytes, _ := json.MarshalIndent(webHook, " ", "  ")

		log.Println("Request payload:")
		log.Println(string(bytes))
	}
}

func logWebHookIfDebug(hook *WebHook) {
	if Config.Debug {
		jsonData, _ := json.MarshalIndent(hook, " ", " ")

		log.Println("Parsed request webhook as json:")
		log.Printf("%s\n", jsonData)
	}
}
