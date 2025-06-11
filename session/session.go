package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func SignUp(client *http.Client, url *url.URL, email string, password string) {
	marshaledBody, err := json.Marshal(CredentialsRequestBody{Email: email, Password: password})
	if err != nil {
		log.Printf("\nError marshaling request body: %v\n", err.Error())
		return
	}

	r, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(marshaledBody))
	if err != nil {
		log.Printf("\nFailed to create a post request: %v\n", err)
		return
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		log.Printf("\nFailed to perform a post request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("\nFailed to read the response body: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusCreated {
		fmt.Printf("Request failed.\nStatus Code: %v\nBody: %v\n", res.StatusCode, string(body))
	}
}

func LogIn(client *http.Client, url *url.URL, email string, password string) {
	marshaledBody, err := json.Marshal(CredentialsRequestBody{Email: email, Password: password})
	if err != nil {
		log.Printf("\nError marshaling request body: %v\n", err.Error())
		return
	}

	r, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(marshaledBody))
	if err != nil {
		log.Printf("\nFailed to create a post request: %v\n", err)
		return
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		log.Printf("\nFailed to perform a post request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("\nFailed to read the response body: %v\n", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed.\nStatus Code: %v\nBody: %v\n", res.StatusCode, string(body))
	}

	fmt.Println(len(client.Jar.Cookies(url)))
}

func LogOut(client *http.Client, url *url.URL) {
	r, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		log.Printf("\nFailed to create a post request: %v\n", err)
		return
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		log.Printf("\nFailed to perform a post request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("\nFailed to read the response body: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed.\nStatus Code: %v\nBody: %v\n", res.StatusCode, string(body))
	}
}

func DropOut(client *http.Client, url *url.URL) {
	r, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		log.Printf("\nFailed to create a post request: %v\n", err)
		return
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		log.Printf("\nFailed to perform a post request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("\nFailed to read the response body: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed.\nStatus Code: %v\nBody: %v\n", res.StatusCode, string(body))
	}
}
