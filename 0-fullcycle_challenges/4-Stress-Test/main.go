package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Uso: --url=<URL> --requests=<número> --concurrency=<número>")
		return
	}

	start := time.Now()

	results := make(chan Result, *requests)
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, *concurrency)

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}        // acquire
			defer func() { <-semaphore }() // release

			reqStart := time.Now()
			resp, err := http.Get(*url)
			duration := time.Since(reqStart)

			if err != nil {
				results <- Result{StatusCode: 0, Duration: duration}
				return
			}
			defer resp.Body.Close()

			results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		}()
	}

	wg.Wait()
	close(results)

	totalTime := time.Since(start)

	// Coletar métricas
	statusCounts := make(map[int]int)
	var totalRequests int
	var status200 int

	for result := range results {
		totalRequests++
		statusCounts[result.StatusCode]++
		if result.StatusCode == 200 {
			status200++
		}
	}

	// Relatório
	fmt.Printf("Tempo total gasto na execução: %v\n", totalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", status200)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCounts {
		if code != 200 {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}
