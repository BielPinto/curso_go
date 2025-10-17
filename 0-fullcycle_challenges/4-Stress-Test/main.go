package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"
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

	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de perfil CPU: %v\n", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("Erro ao iniciar perfil CPU: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

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

	// Stop CPU profiling
	pprof.StopCPUProfile()

	// Write memory profile
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de perfil de memória: %v\n", err)
		return
	}
	defer memFile.Close()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Printf("Erro ao escrever perfil de memória: %v\n", err)
		return
	}

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

	// Criar arquivo de relatório
	reportFile, err := os.Create("report.txt")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de relatório: %v\n", err)
		return
	}
	defer reportFile.Close()

	// Escrever relatório no arquivo
	fmt.Fprintf(reportFile, "Tempo total gasto na execução: %v\n", totalTime)
	fmt.Fprintf(reportFile, "Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Fprintf(reportFile, "Quantidade de requests com status HTTP 200: %d\n", status200)
	fmt.Fprintf(reportFile, "Distribuição de outros códigos de status HTTP:\n")
	for code, count := range statusCounts {
		if code != 200 {
			fmt.Fprintf(reportFile, "  %d: %d\n", code, count)
		}
	}

	fmt.Println("Relatório salvo em report.txt")
	fmt.Println("Perfis salvos: cpu.prof e mem.prof")
}
