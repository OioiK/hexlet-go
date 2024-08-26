package main

import (
	"fmt"

	greetingv1 "./greeting"
	greetingv2 "./greeting/v2"
)

func main() {
	fmt.Println("Первое приветствие: ", greetingv1.Get(), "\n", "Второе приветствие: ", greetingv2.Get())
}
