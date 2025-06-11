package main

import (
	"github.com/joho/godotenv"
)

var a *App

func init() {
	_ = godotenv.Load()

	a = &App{}
	a.Initialize()
}

func main() {
	handleMainMenu(a)
}
