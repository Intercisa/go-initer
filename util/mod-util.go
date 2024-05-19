package util

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetMain() string {
	return `
	package main

	import(
		"fmt"
	)

	func main() {
		fmt.Println("Hello modul!")
	}

	`
}
