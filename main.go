package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getChoice(reader *bufio.Reader) string {
	isPrefix := true
	var line, partialLine []byte
	var err error
	for isPrefix {
		partialLine, isPrefix, err = reader.ReadLine()

		if err != nil {
			log.Fatal("Failed to read from stdin.")
		}
		line = append(line, partialLine...)
	}

	return string(line[0])
}

func main() {
	_ = godotenv.Load()

	a := &App{}
	a.Initialize()
	handleMainMenu(a)
}

func handleMainMenu(a *App) {
	for {
		fmt.Printf(`%vMain
---------
1 Session Token
2 JSON Web Token
3 OAuth
4 Exit
---------
> `, "\n")

		switch string(getChoice(a.Reader())) {
		case "1":
			handleTokenMenu(a, "Session Token")
			continue
		case "2":
			handleTokenMenu(a, "JSON Web Token")
			continue
		case "3":
		case "4":
			fmt.Println("Exiting.")
			os.Exit(0)

		default:
			fmt.Println("Invalid option")
			continue
		}
	}
}

func handleTokenMenu(a *App, tokenType string) {
	for {
		fmt.Printf(`%v%v
-----------------------
1 Sign up
2 Log in
3 Log out
4 Back (implies logout)
-----------------------
> `, "\n", tokenType)

		switch getChoice(a.Reader()) {
		case "1":
			{
				url := fmt.Sprintf("http://%v:%v/session/signup", a.ServerHost(), a.ServerPort())
				body := []byte(`
				{
				"username": "shr",
				"password": "esw"
				}
				`)
				r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
				if err != nil {
					log.Printf("\nFailed to create a post request: %v\n", err)
					continue
				}
				r.Header.Add("Content-Type", "application/json")
				res, err := http.DefaultClient.Do(r)
				if err != nil {
					log.Printf("\nFailed to perform a post request: %v\n", err)
					continue
				}
				defer res.Body.Close()
				fmt.Println(res.Body)
				continue
			}

		case "2":
		case "3":
		case "4":
			return

		default:
			fmt.Println("Invalid option")
			continue
		}
	}
}
