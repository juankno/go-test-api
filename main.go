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

	for _, api := range apis {
		checkApi(api)
	}

	elapsed := time.Since(startTime)

	fmt.Printf("¡Listo, tomó %v segundos!\n", elapsed.Seconds())
}

func checkApi(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("¡Error: %s esta caído!\n", api)
		return
	}
	fmt.Printf("¡Success: %s está en funcionamiento!\n", api)
}
