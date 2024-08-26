package main

import (
	color "github.com/fatih/color"
	"github.com/sirupsen/logrus"
) // Указываем путь до нужного пакета внутри репозитория

func main() {
	logrus.Println("Hello, Hexlet!")
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Hello, Hexlet!")
}
