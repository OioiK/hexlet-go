package main

import (
	"hexlet-go/greeting"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
) // Указываем путь до нужного пакета внутри репозитория

func main() {
	logrus.Println(greeting.Get())
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println(greeting.Get())
}
