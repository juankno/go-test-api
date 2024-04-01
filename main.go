package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	startTime := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereinthernet.com",
		"https://graph.microsoft.com",
	}

	channel := make(chan string)

	for _, api := range apis {
		go checkApi(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-channel)
	}

	elapsed := time.Since(startTime)

	fmt.Printf("¡Listo, tomó %v segundos!\n", elapsed.Seconds())
}

func checkApi(api string, channel chan string) {
	if _, err := http.Get(api); err != nil {
		channel <- fmt.Sprintf("¡Error: %s esta caído!\n", api)
		return
	}
	channel <- fmt.Sprintf("¡Success: %s está en funcionamiento!\n", api)
}
