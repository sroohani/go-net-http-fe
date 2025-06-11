package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/sroohani/go-net-http-fe/session"
)

const trimChars = "\a\b\t\n\f\r\v "

type AuthType = int

const (
	AuthSessionToken AuthType = iota
	AuthJsonWebToken
	AuthOAuth
)

func getMenuDetails(authType AuthType) (string, string) { // (Menu title, Endpoint)
	switch authType {
	case AuthSessionToken:
		return "Session Token", "session"
	case AuthJsonWebToken:
		return "JSON Web Token", "jwt"
	case AuthOAuth:
		return "OAuth", "oauth"
	default:
		return "", ""
	}
}

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

	return string(strings.Trim(string(line), trimChars)[0])
}

func getEmail(reader *bufio.Reader) string {
	fmt.Println("Please enter your email address:")
	fmt.Println("(If you wish to cancel, press <SPACE> and then <ENTER>)")
	fmt.Print("> ")

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

	return string(strings.Trim(string(line), trimChars))
}

func getPassword() string {
	fmt.Println("Please enter your password:")
	fmt.Println("(If you wish to cancel, press <SPACE> and then <ENTER>)")
	fmt.Print("> ")

	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Failed to read from stdin.")
	}

	return string(strings.Trim(string(password), trimChars))
}

func handleMainMenu(a *App) {
	for {
		fmt.Printf(`%vMain
----------------
1 Session Token
2 JSON Web Token
3 OAuth
4 Exit
----------------
> `, "\n")

		switch string(getChoice(a.Reader())) {
		case "1":
			handleTokenMenu(a, AuthSessionToken)
			continue
		case "2":
			handleTokenMenu(a, AuthJsonWebToken)
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

func handleTokenMenu(a *App, authType AuthType) {
	menuTitle, endpoint := getMenuDetails(authType)

	for {
		fmt.Printf(`%v%v
-----------------------
1 Sign up
2 Log in
3 Log out
4 Drop out
5 Back (implies logout)
-----------------------
> `, "\n", menuTitle)

		switch getChoice(a.Reader()) {
		case "1":
			email := getEmail(a.Reader())
			if email == "" {
				fmt.Println("Canceled by the user.")
				continue
			}
			password := getPassword()
			if password == "" {
				fmt.Println("Canceled by the user.")
				continue
			}
			fmt.Println()
			session.SignUp(a.client, a.url.JoinPath(fmt.Sprintf("%v/signup", endpoint)), email, password)
			continue

		case "2":
			email := getEmail(a.Reader())
			if email == "" {
				fmt.Println("Canceled by the user.")
				continue
			}
			password := getPassword()
			if password == "" {
				fmt.Println("Canceled by the user.")
				continue
			}
			fmt.Println()
			session.LogIn(a.client, a.url.JoinPath(fmt.Sprintf("%v/login", endpoint)), email, password)
			continue

		case "3":
			fmt.Println()
			session.LogOut(a.client, a.url.JoinPath(fmt.Sprintf("%v/logout", endpoint)))
			continue

		case "4":
			fmt.Println()
			session.DropOut(a.client, a.url.JoinPath(fmt.Sprintf("%v/dropout", endpoint)))
			continue

		case "5":
			return

		default:
			fmt.Println("Invalid option")
			continue
		}
	}
}
