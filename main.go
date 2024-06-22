package main

import (
	"fmt"
	"net/http"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(apis))

	results := make(chan string, len(apis))

	for _, api := range apis {
		go checkApi(api, results, &wg)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Print(result)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("¡Listo, tomó %v segundos!\n", elapsed.Seconds())
}

func checkApi(api string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(api)
	if err != nil {
		results <- fmt.Sprintf("¡Error: %s está caído! Detalle: %v\n", api, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Sprintf("¡Error: %s devolvió estado %d!\n", api, resp.StatusCode)
		return
	}

	results <- fmt.Sprintf("¡Success: %s está en funcionamiento!\n", api)
}
