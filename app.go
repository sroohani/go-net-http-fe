package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
)

type App struct {
	serverHost string
	serverPort int
	client     *http.Client
	jar        *cookiejar.Jar
	url        *url.URL
	reader     *bufio.Reader
}

func (a *App) Initialize() {
	a.serverHost = os.Getenv("SERVER_HOST")
	if a.serverHost == "" {
		a.serverHost = "localhost"
	}
	serverPortStr := os.Getenv("SERVER_HOST")
	if serverPortStr == "" {
		serverPortStr = "9876"
	}
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		serverPort = 9876
	}
	a.serverPort = serverPort

	a.url = &url.URL{Scheme: "http", Host: fmt.Sprintf("%v:%v", a.serverHost, a.serverPort), Path: "/"}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("\nError creating cookie jar: %v\n", err)
	}
	a.jar = jar
	a.client = &http.Client{
		Jar: jar,
	}

	a.reader = bufio.NewReader(os.Stdin)
}

func (a *App) ServerHost() string {
	return a.serverHost
}

func (a *App) ServerPort() int {
	return a.serverPort
}

func (a *App) Reader() *bufio.Reader {
	return a.reader
}
